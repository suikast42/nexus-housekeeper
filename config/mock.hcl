loglevel= "debug"

nexusServer{
  scheme = "https"
  host = "nexus.cloud.private"
  username= "admin"
  password= "cloudmaster"
  keepImages {
    default = 1
    dockerLocal=1
  }
}
