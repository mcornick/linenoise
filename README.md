# linenoise

__linenoise__ is a library that generates strings of random characters (herein called a "noise") that can be used as reasonably secure passwords.It is an extraction of the password generator from my prior project [genpw](https://github.com/mcornick/genpw), and intended to be genpw's successor.

## Interface

__linenoise__ exports one function and one struct.

`Noise` is the noise-generating function. It is called with a `Parameters` (see next paragraph.) It returns a `string` and an `error`.  If the `Parameters` can be used to create a noise with the desired length, the `string` is the generated noise and the `error` is `nil`.  Otherwise, the `string` is `""` and the `error` is a typical Go error object.

`Parameters` is a struct containing the following:

* `Length int` is the length of the noise desired.
* `Upper bool` indicates whether the noise should contain uppercase characters.
* `Lower bool` indicates whether the noise should contain lowercase characters.
* `Digit bool` indicates whether the noise should contain digits.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/mcornick/linenoise.

This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](https://www.contributor-covenant.org/) code of conduct.

## License

__linenoise__ is available as open source under the terms of the MIT License.
