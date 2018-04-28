# udebug-cli

Command line interface for udebug platform.


## Set credentials

Create .env file and put the following values inside:
```
UDEBUG_USERNAME=<username>
UDEBUG_PASSWORD=<password>
```

Source the credentials

```
export $(cat .env | grep -v ^# | xargs)
```

## Explore

```
udebug-cli --help
```

## NOTE

Not all sentinels are inplace yet, so beware of positional arguments. =)