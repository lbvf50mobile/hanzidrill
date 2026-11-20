# Hanzidrill - Stroke Order Animation and Prononciation


https://lbvf50mobile.github.io/hanzidrill/

This tool purpose is to start writing new Hanzi characters, the first
approach. Getting animation and sound. Helps to make first writes of Hanzi
Charackters and Hanzi Charackters combitations.

First it was a very straightforwad solution based on one HTTP server writen in
Golgand and set of static files. Further migrated to the Jekyll static
generator to split hardcoded HTML into parts and simplify process of new
characters additions. 

## Oh how could I start it local?

Learning HARD WAY - the GitHub looks like to read config from the
`docs/_config.yml` no need to create all these additional configs.


## Just if you have nothing to do and wan to make some experiments. 

I definatelly think you could. But there is a caveat. Different configurations
for different hosts, different config files so need to do somethigh like that:

```bash
alias shanzi="jekyll serve --host ${MY_IP} -P 3888 --livereload  --config _config.yml,_config_local.yml"
```
So need to add `export MY_IP=192.168.1.5` or whatever somewhere in
`.bashrc`. I just want to sleep and have to move this project and fix it from
the Bugs I never intend to do. So Here it is.

## I run several instances and have a glictch? What should I do?

Yep, body. That's all because `--livereload` so one instance with
`--livereload` and other without. Let the `.bashrc` grow, more aliaces.
