provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_cloud_run_service" "order-service" {
  name     = "order-service"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/kibandaa-236d4/order-service:latest"
      }
    }
  }
}

resource "google_cloud_run_service" "payment-service" {
  name     = "payment-service"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/kibandaa-236d4/payment-service:latest"
      }
    }
  }
}
