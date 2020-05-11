# How to make the Google form

1. Create a new Google form.
2. Make the question type "Checkbox grid".
3. Click the "..." next to "Require a response in each row" and
   click on "Limit to one response per column". This will ensure that
   the same ranking can't be used twice.
4. Enter candidate names under "Rows".
5. Enter rankings under "Columns".


# Usage

First, download a CSV of the responses to your Google form by going to "Responses",
then clicking on the "..." and selecting "Download responses (.csv)".

Then run the tabulator on the file:

```
# With `go run`
$ go run *.go --possible-rankings COMMA_SEPARATED_RANKINGS --csv-file RESPONSES_CSV_FILE
```

If the possible rankings are `1`, `2`, and `3`, for example, then the value of
`COMMA_SEPARATED_RANKINGS` should be `1,2,3`.
