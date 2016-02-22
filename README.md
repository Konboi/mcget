# mcget

mcget is mirage list api parser.

mcget can get mirage docker container id from input subdomain.

# What is mirage

[mirage](https://github.com/acidlemon/mirage) is docker front end & reverse proxy for development

# Useage

```
$(mcget -h <mirage host uri ex) http://mirage.baseurl.com> -s <subdomain> bash)
# $ docker exec -it <sub domain docker container id> bash
```
