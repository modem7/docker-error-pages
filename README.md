<p align="center">
  <img src="https://hsto.org/webt/rm/9y/ww/rm9ywwx3gjv9agwkcmllhsuyo7k.png" width="94" alt="" />
</p>

# HTTP's error pages in Docker image

[![Build Status](https://drone.modem7.com/api/badges/modem7/docker-error-pages/status.svg)](https://drone.modem7.com/modem7/docker-error-pages)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/modem7/traefik-error-pages/latest)
![Docker Pulls](https://img.shields.io/docker/pulls/modem7/traefik-error-pages)
![License](https://img.shields.io/github/license/modem7/docker-error-pages)

This repository contains:

- A very simple [generator](ErrorPages/generator/generator.js) _(`nodejs`)_ for HTTP error pages _(like `404: Not found`)_ with different templates supports
- Dockerfile for docker image ([docker hub][link_docker_hub], [ghcr.io][link_ghcr]) with generated pages and `nginx` as a web server

**Can be used for [Traefik error pages customization](https://doc.traefik.io/traefik/middlewares/http/errorpages/)**.

## Demo

Generated pages (from the latest release) always **[accessible here](https://modem7.github.io/docker-error-pages/)** _(live preview)_.

## Templates

Name              | Preview
:---------------: | :-----:
`ghost`           | [![ghost](https://hsto.org/webt/oj/cl/4k/ojcl4ko_cvusy5xuki6efffzsyo.gif)](https://tarampampam.github.io/error-pages/ghost/404.html)
`l7-light`        | [![l7-light](https://hsto.org/webt/xc/iq/vt/xciqvty-aoj-rchfarsjhutpjny.png)](https://tarampampam.github.io/error-pages/l7-light/404.html)
`l7-dark`         | [![l7-dark](https://hsto.org/webt/s1/ih/yr/s1ihyrqs_y-sgraoimfhk6ypney.png)](https://tarampampam.github.io/error-pages/l7-dark/404.html)
`shuffle`         | [![shuffle](https://hsto.org/webt/7w/rk/3m/7wrk3mrzz3y8qfqwovmuvacu-bs.gif)](https://tarampampam.github.io/error-pages/shuffle/404.html)
`hacker-terminal` | [![hacker-terminal](https://hsto.org/webt/5s/l0/p1/5sl0p1_ud_nalzjzsj5slz6dfda.gif)](https://tarampampam.github.io/error-pages/hacker-terminal/404.html)
`hexxone`         | [Hexxone](https://github.com/hexxone/error-pages)
`parallax`        | [Thom-x](https://github.com/Thom-x/docker-error-pages)

## Usage

Generated error pages in our [docker image][link_docker_hub] permanently located in directory `/opt/html/%TEMPLATE_NAME%`. `nginx` in a container listen for `8080` (`http`) port.

#### Supported environment variables

Name            | Description
--------------- | -----------
`TEMPLATE_NAME` | (`ghost` by default) "default" pages template _(allows to use error pages without passing theme name in URL - `http://127.0.0.1/500.html` instead `http://127.0.0.1/ghost/500.html`)_

Also, you can use a special template name `random` - in this case template will be selected randomly.

### Ready docker image

[![dockeri.co](https://dockeri.co/image/modem7/traefik-error-pages)](https://hub.docker.com/r/modem7/traefik-error-pages)

Execute in your shell:

```bash
$ docker run --rm -p "8082:8080" modem7/error-pages:latest
```

And open in your browser `http://127.0.0.1:8082/ghost/400.html`.

### Custom error pages for [Traefik][link_traefik]

Simple traefik (tested on `v2.5`) service configuration (**change with your needs**):

```yaml
version: '3.4'

services:
  error-pages:
    image: modem7/traefik-error-pages:latest
    environment:
      TEMPLATE_NAME: l7-dark
    networks:
      - traefik-public
    deploy:
      placement:
        constraints:
          - node.role == worker
      labels:
        traefik.enable: 'true'
        traefik.docker.network: traefik-public
        # use as "fallback" for any non-registered services (with priority below normal)
        traefik.http.routers.error-pages-router.rule: HostRegexp(`{host:.+}`)
        traefik.http.routers.error-pages-router.priority: 10
        # should say that all of your services work on https
        traefik.http.routers.error-pages-router.tls: 'true'
        traefik.http.routers.error-pages-router.entrypoints: https
        traefik.http.routers.error-pages-router.middlewares: error-pages-middleware@docker
        traefik.http.services.error-pages-service.loadbalancer.server.port: 8080
        # "errors" middleware settings
        traefik.http.middlewares.error-pages-middleware.errors.status: 400-599
        traefik.http.middlewares.error-pages-middleware.errors.service: error-pages-service@docker
        traefik.http.middlewares.error-pages-middleware.errors.query: /{status}.html

  any-another-http-service:
    image: nginx:alpine
    networks:
      - traefik-public
    deploy:
      placement:
        constraints:
          - node.role == worker
      labels:
        traefik.enable: 'true'
        traefik.docker.network: traefik-public
        traefik.http.routers.another-service.rule: Host(`subdomain.example.com`)
        traefik.http.routers.another-service.tls: 'true'
        traefik.http.routers.another-service.entrypoints: https
        # next line is important
        traefik.http.routers.another-service.middlewares: error-pages-middleware@docker
        traefik.http.services.another-service.loadbalancer.server.port: 80

networks:
  traefik-public:
    external: true
```

## License

This is open-sourced software licensed under the [MIT License](https://github.com/modem7/docker-error-pages/blob/master/LICENSE).
