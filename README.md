# Element Repository Stats

Want to get pretty dashboards like this? Read on!

<img width="837" alt="Screenshot 2022-12-08 at 19 34 52" src="https://user-images.githubusercontent.com/7190048/206551922-68cbbc5b-4297-4405-9c44-5a8cdca41255.png">

<img width="1690" alt="Screenshot 2022-12-09 at 17 13 00" src="https://user-images.githubusercontent.com/7190048/206756683-462bb5b4-6b6d-482d-bff2-c4a383680c37.png">


## Requirements

You need Docker.

Eligible repositories must:
- live in Github.
- use the [labelling scheme](https://github.com/vector-im/element-meta/wiki/Triage-process).

You need to:
- create a [personal access token](https://github.com/settings/tokens) (Classic, not fine-grained) for Github with `repo:status`, `public_repo`, `user:email` scopes. You may need more if you want to run this in private repositories or see members who have not publicly said they are inside a given organisation. Copy the token into a file called `gh-token` at the top-level of this repository.
- state which repositories you are interested in. Currently this is done by [modifying the docker-compose.yml file](https://github.com/kegsay/github_exporter/blob/master/docker-compose.yml#L34).

## Running

```
docker-compose up
```

Then visit http://localhost:3000/d/UGK4RFiGk/github?orgId=1&refresh=1m&from=now-5m&to=now - you'll need to wait several minutes for data to show up. You can check some graphs almost immediately though.

## How do I get pull request latency for a certain time period?

- Find out which PR number range you need e.g #28xx-#29xx
- Edit the PR latency pie chart to filter to these PRs:

<img width="1383" alt="Screenshot 2022-12-09 at 16 57 16" src="https://user-images.githubusercontent.com/7190048/206756751-23658876-aedd-41a3-9cf0-11c0860fb2fb.png">


## How does this work?

- Grafana is a pretty visualiser. It reads data from Prometheus.
- Prometheus is a glorified metrics scraper and database. It scrapes data from exporters.
- Github exporter (aka THIS PROJECT) is an exporter which exposes github data for repositories. It's based on https://github.com/xrstf/github_exporter 
- Docker-compose just ties this all together into a single command.
