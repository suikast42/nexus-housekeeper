## Generate nexus client from openapi doc
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate  -i http://172.21.1.10:5002/service/rest/swagger.json -l go -o /local/client/nexus3 --additional-properties packageName=nexus3
