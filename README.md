`aws sts assume-role` returns a json object containing the credentials in order to assume the requested role.

pipe this output into this tool and it will provide you with the `export ENV_VAR` statements to set your shell to use the requested role.

## Example

```
$ aws sts assume-role --role-arn arn:aws:iam::123456789123:role/requested-role --role-session-name a-session-name | sts-assume-role-env-vars
export AWS_ACCESS_KEY_ID=ASIA....
export AWS_SECRET_ACCESS_KEY=lakshjdjkahsd
export AWS_SESSION_TOKEN=askljdakljsd
```
or if you are even lazier than that
```
$ eval $(aws sts assume-role --role-arn arn:aws:iam::123456789123:role/requested-role --role-session-name a-session-name | sts-assume-role-env-vars)
```

just copy and paste these and you'll be using the requested session / role.

