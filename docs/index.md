---
title: linenoise
---
__linenoise__ is a library that generates strings of random characters (herein called a "noise") that can be used as reasonably secure passwords.It is an extraction of the password generator from my prior project _genpw_, and intended to be genpw's successor.

__linenoise__ 1.0.x and 1.1.x are stable. There will be no new features added to __linenoise__ 1.x.x. New features may be added in a 2.0.x branch, if I think of any, but I have none planned right now.

## Interface

__linenoise__ exports one function and one struct.

`Noise` is the noise-generating function. It is called with a `Parameters` (see next paragraph.) It returns a `string` and an `error`.  If the `Parameters` can be used to create a noise with the desired length, the `string` is the generated noise and the `error` is `nil`.  Otherwise, the `string` is `""` and the `error` is a typical Go error object.

`Parameters` is a struct containing the following:

* `Length int` is the length of the noise desired.
* `Upper bool` indicates whether the noise should contain uppercase characters.
* `Lower bool` indicates whether the noise should contain lowercase characters.
* `Digit bool` indicates whether the noise should contain digits.

## Example

```go
import "github.com/mcornick/linenoise"

p := linenoise.Parameters{
        Length: 42,
        Upper:  true,
        Lower:  false,
        Digit:  true,
}
result, err := linenoise.Noise(p)
if err != nil {
        log.Fatal(err)
}
fmt.Println(result)
```

## Contributing to linenoise

If you think you have a problem, improvement, or other contribution towards the betterment of __linenoise__, please file an issue or, where appropriate, a pull request.

Keep in mind that I'm not paid to write Go code, so I'm doing this in my spare time, which means it might take me a while to respond.

When filing a pull request, please explain what you're changing and why. Please use standard Go formatting (`go fmt` is your friend.) Please limit your changes to the specific thing you're fixing and isolate your changes in a topic branch that I can merge without pulling in other stuff.

__linenoise__ uses [Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog-angular/blob/master/convention.md) style. Please follow this convention. Scopes are not required in commit messages.

__linenoise__ uses the MIT license. Please indicate your acceptance of the MIT license by using [git commit --signoff](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt--s).

__linenoise__ is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

Thanks for contributing!
