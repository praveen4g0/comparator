# Validate Comparator library

This is an executable specification file. This file follows markdown syntax.
Every heading in this file denotes a scenario. Every bulleted point denotes a step.

To execute this specification, run

    gauge specs


## Compare API Responses from files
tags: e2e

* Compare API response
  |file-dir1         |file-dir2         |
  |------------------|------------------|
  |testdata/json1.txt|testdata/json2.txt|

## Compare API Response from files (with No equal data)
tags: e2e

* Compare API response
  |file-dir1         |file-dir2         |
  |------------------|------------------|
  |testdata/json1.txt|testdata/json3.txt|


## Compare API Response with invalid url in file 2
tags: e2e

* Compare API response
  |file-dir1         |file-dir2         |
  |------------------|------------------|
  |testdata/json3.txt|testdata/json4.txt|

  ## Compare API Response with invalid url in file 1
tags: e2e

* Compare API response
  |file-dir1         |file-dir2         |
  |------------------|------------------|
  |testdata/json4.txt|testdata/json3.txt|
