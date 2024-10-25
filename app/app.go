package app

import (
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/suikast42/nexus-housekeeper/client/nexus3"
	"net/http"
	"sort"
	"strings"
)

func (a *App) DeleteOldContainers() error {
	ofType, err := a.getReposOfType("docker")

	if err != nil {
		return err
	}
	for _, repo := range ofType {
		a.Logger().Info().Msg(fmt.Sprintf("%s", repo.Name))
		assets, err := a.getImagesOfRepo(repo)
		if err != nil {
			return err
		}
		for k, components := range a.nexusComponentsMap(assets) {
			val, ok := a.Config().NexusServer.KeepImages[repo.Name]
			if !ok {
				val = a.Config().NexusServer.KeepImages["default"]
			}
			sortVersions(components)
			if val > 0 && len(components) > val {
				a.Logger().Info().Msg(fmt.Sprintf("\tDeleteing Versions of %s", k))
				// Starting to iterate from the first delete pos
				for _, component := range components[val:] {
					a.Logger().Info().Msg(fmt.Sprintf("\t\t%s", component.Version))
					err := a.deleteComponent(component)
					if err != nil {
						a.Logger().Err(err)
					}
					//for _, asset := range component.Assets {
					//	a.Logger().Info().Msg(fmt.Sprintf("\t\t\t%s", asset.Path))
					//}
				}
			}

		}
		//for _, asset := range a.nexusComponentsMap(assets) {
		//	a.Logger().Info().Msg(fmt.Sprintf("\t%s %s", asset.Name, asset.Version))
		//}
	}
	return nil
}

func (a *App) deleteComponent(xo nexus3.ComponentXo) error {
	response, err := a.ClientNexus().ComponentsApi.DeleteComponent(a.Context(), xo.Id)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return errors.New(statusCodeToString(response))
	}

	return nil
}
func (a *App) getReposOfType(_type string) ([]nexus3.AbstractApiRepository, error) {
	var result = []nexus3.AbstractApiRepository{}
	repos, response, err := a.ClientNexus().RepositoryManagementApi.GetRepositories(a.Context())
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(statusCodeToString(response))
	}
	for _, repo := range repos {
		if strings.EqualFold(repo.Format, _type) {
			result = append(result, repo)
		}
	}

	return result, nil
}

func (a *App) getImagesOfRepo(repo nexus3.AbstractApiRepository) ([]nexus3.ComponentXo, error) {
	components, response, err := a.ClientNexus().ComponentsApi.GetComponents(a.Context(), repo.Name, &nexus3.ComponentsApiGetComponentsOpts{})
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(statusCodeToString(response))
	}
	return components.Items, nil
}

func (a *App) nexusComponentsMap(components []nexus3.ComponentXo) map[string][]nexus3.ComponentXo {
	result := make(map[string][]nexus3.ComponentXo)
	for _, compo := range components {
		val, ok := result[compo.Name]
		if ok {
			compo2 := append(val, compo)
			result[compo.Name] = compo2
		} else {
			result[compo.Name] = []nexus3.ComponentXo{compo}
		}
	}
	return result
}
func statusCodeToString(response *http.Response) string {
	return fmt.Sprintf("Status code %d. Status message: %s", response.StatusCode, response.Status)
}

// sortVersions Desc in Version or LastModified
func sortVersions(compos []nexus3.ComponentXo) {
	sort.Slice(compos, func(i, j int) bool {
		v1, err1 := semver.ParseTolerant(compos[i].Version)
		v2, err2 := semver.ParseTolerant(compos[j].Version)

		//In case of parsing errors we're using the ts of the asset
		if err1 != nil || err2 != nil {
			if len(compos[i].Assets) > 0 && len(compos[j].Assets) > 0 {
				return compos[i].Assets[len(compos[i].Assets)-1].LastModified.After(compos[j].Assets[len(compos[j].Assets)-1].LastModified)
			}
			return false
		}

		return v1.GT(v2)
	})

}
