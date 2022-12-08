# Element Repository Stats

Want to get pretty dashboards like this? Read on!

<img width="837" alt="Screenshot 2022-12-08 at 19 34 52" src="https://user-images.githubusercontent.com/7190048/206551922-68cbbc5b-4297-4405-9c44-5a8cdca41255.png">



## Requirements

You need Docker.

Eligible repositories must:
- live in Github.
- use the [labelling scheme](https://github.com/vector-im/element-meta/wiki/Triage-process).

You need to:
- create a [personal access token](https://github.com/settings/tokens) (Classic, not fine-grained) for Github with `repo:status`, `public_repo`, `user:email` scopes. You may need more if you want to run this in private repositories or see members who have not publicly said they are inside a given organisation. Copy the token into a file called `gh-token` at the top-level of this repository.
- state which repositories you are interested in. Currently this is done by modifying the docker-compose.yml file and defaults to `matrix-org/dendrite`.
- Build the exporter docker image: `docker build -t github-exporter .`

## Running

```
docker-compose up
```

Wait a few minutes for everything to be populated. Then visit http://localhost:3000/d/UGK4RFiGk/github?orgId=1&refresh=1m&from=now-5m&to=now 

## How does this work?

- Grafana is a pretty visualiser. It reads data from Prometheus.
- Prometheus is a glorified metrics scraper and database. It scrapes data from exporters.
- Github exporter (aka THIS PROJECT) is an exporter which exposes github data for repositories. It's based on https://github.com/xrstf/github_exporter 
- Docker-compose just ties this all together into a single command.
