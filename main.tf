resource "docker_image" "golang" {
  name         = "golang:1.25"  # Go 1.25 version
  keep_locally = true
}

resource "docker_container" "app" {
  name  = "devops-security-app"
  image = docker_image.golang.name
  tty   = true

  volumes {
    host_path      = "C:/Users/PIS/devops-security-platform"  # <-- absolute path for Windows
    container_path = "/workspace"
  }

  working_dir = "/workspace"
  command     = ["sleep", "infinity"]
}