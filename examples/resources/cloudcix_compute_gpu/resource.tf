resource "cloudcix_compute_gpu" "example" {
  project_id  = cloudcix_project.example.id
  instance_id = cloudcix_compute_instance.example.id
  name        = "gpu01"

  specs = [{
    sku_name = "A100_GPU"
  }]
}
