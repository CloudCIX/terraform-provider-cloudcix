resource "cloudcix_storage_volume" "example_storage_volume" {
  project_id = 1
  specs = [{
    quantity = 250
    sku_name = "SSD_001"
  }]
  instance_id = 456
  name = "additional-storage"
  type = "hyperv"
}
