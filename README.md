# R.E.D Copy

Redundant Error-resistant Digital Copy is a set of functions and a an accompanying CLI tool which allow for the creation and reading of error redundant files by multiplying the occurences of bits. 


## Use

The redcopy CLI can be installed and used as a stand-alone tool, or the functionallity can be incorporated into your own projects with the use of the `redlib`.

### CLI 

The CLI uses sub-commands, such as `redcopy read [read args]` and `redcopy write [write args]` each with their own positional arguments.

#### `redcopy write`

`redcopy write` takes a file and generates a RED file with multiplied bits. It takes the follwoing positional arguments:
`[source file path] [multiplicity] [header]`

Multiplicity is a preset among the following. Omitting a multiplicity will default to `normal`.

`"weak": mult = 3`
`"normal": mult = 8`
`"strong": mult = 32`
`"super": mult = 64`
`"ludicrous": mult = 254`

Header places a single byte at the head of the file, the value of which is the multiplicity. Defaults to true, can be set to false to skip the header. *Note: these are positional arguments! If you want to ovewrite the default header toggle, you must specify a multiplicity.*


#### `redcopy read`

`redcopy read` will demultiply a RED file into its original bits. It takes the following positional arguments:

`[source file path] [output file path] [multiplicity]`

If the RED file has a header, the header will automatically be used as multiplicity and the third argument is not needed.

If the RED file does not have a header the thir argument must be one of the five multiplicity presets:

`"weak": mult = 3`
`"normal": mult = 8`
`"strong": mult = 32`
`"super": mult = 64`
`"ludicrous": mult = 254`




### TODOs

* Test Suite
* Corruption aware multiplicity guessing
* Corruption aware bit devide
