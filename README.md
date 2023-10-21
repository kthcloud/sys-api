# 💻 kthcloud/sys-api
[![ci](https://github.com/kthcloud/sys-api/actions/workflows/docker-image.yml/badge.svg)](https://github.com/kthcloud/sys-api/actions/workflows/docker-image.yml)

This API is used to fetch capacities, usage and other information about the kthcloud system.

It does this by fetching information from every host
(using the [host-api](https://github.com/kthcloud/host-api) installed locally) and Kubernetes clusters

## 📚 Docs
This API is documented using Swagger.

You can view the publically hosted documentation [here](https://api.cloud.cbh.kth.se/landing/v1/docs/index.html).

## 🤝 Contributing

Thank you for considering contributing!

Right now we don't support running the API outside kthcloud,
so you would need to set up a Kubernetes cluster, CloudStack environment, Harbor service and a GitLab instance.

If you are part of the development team at kthcloud,
you will find the current configuration file with the required secrets in our private admin repository.

## 📝 License

The API is open-source software licensed under the [MIT license](https://opensource.org/licenses/MIT).

## 📧 Contact

If you have any questions or feedback, submit an Issue!
