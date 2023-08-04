# CSV to XLSX
## Introduction
This command line tool uses a [package](https://github.com/tealeg/csv2xlsx) that converts CSV files to XLSX files 

## Installation

You are going to need the Go programming language installed in order to build the executable.

Once you have Go installed, you'll need to either clone this repository:

#+BEGIN_SRC sh
git clone https://github.com/patch3459/CSVtoXLSX_Converter.git 
#+END_SRC

From within the resulting `CSVtoXLSX_Converter` directory issue the following command to build the project:

```
go build -v .
```
If all goes well you shuould find the compiled binary ```CSVtoXLSX_Converter``` has been created.

## Invocation

This executable takes a directory full of CSV files and converts them all into XLSX files. If you're looking to convert individual files, it may be more convient to use [this](https://github.com/tealeg/csv2xlsx) module instead, which this program is based off of. 

```
./csv2xlsx -f=MyData.csv -o=MyData.xslx
```

If your input file uses a delimiter other than a comma then you must provide that as a third paramater, thus:

```
./csv2xlsx -f="./some_directory" -o="./some_output_directory" -d=";"
```

