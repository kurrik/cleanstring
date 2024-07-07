# cleanstring

Given a string input, cleanstring will:
- Remove any leading blank/whitespace-only lines.
- Strip a prefix consisting of any amount of whitepsace followed by a pipe ("|") character.
- Remove any trailing blank/whitespace-only lines.

The purpose is to allow for clean construction of readable multiline string literals without
including unnecessary whitespace or indentation.

```go
		Steps: []*pb.Step{
			{
				Key: &pb.StepKey{
					Token: ...,
				},
				Script: &pb.UserScript{
					Type: ...,
					Schema: ...,
					Definition: cleanstring.Get(`
						|function step(input: Schema001) {
						|  if (input.state > 0) {
						|    transitionTo2(input);
						|  } else {
						|    input.state *= -1;
						|    transitionTo3(input);
						|  }
						|}
					`),
					Entrypoint: ...,
				},
			},
		}
```

In the example above, `Steps[0].Script.Definition` will be set to:

```typescript
function step(input: Schema001) {
  if (input.state > 0) {
    transitionTo2(input);
  } else {
    input.state *= -1;
    transitionTo3(input);
  }
}
```

## Using

```zsh
go get github.com/kurrik/clearstring
```

Then use:
```go
str := clearstring.Get(`
    |Any literal
    |which needs to be split
    |on multiple lines for readability.
`)
```

`str` will be:
```
Any literal
which needs to be split
on multiple lines for readability.
```