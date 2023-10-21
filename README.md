# Week 5 Assignment: Crawling and Scraping the Web

## Project Summary
This project develops a Go-based web crawler/scraper to obtain text information from a list of target Wikipedia web pages. This project utilizes the [Colly](https://github.com/gocolly/colly) framework to perform the scraping. It then saves page html code to a wikipages directory and creates a JSON lines file (go_items.jl) with each line of JSON representing a Wikipedia page. 

## Files
### *main.go*
This file defines the Scrape() and WriteResults() functions. Scrape takes in a URL string and outputs a struct containing the parsed results from scraping the website as well as the html body. The WriteResults function takes the results from Scrape and writes them to a JSON lines file and to the 'wikipages' directory, respectively. The list of URLs is defined in a slice in the main function which then loops through the slice and calls Scrape and WriteResults.

### *main_test.go*
This file runs a test for the Scrape function. It tests a valid URL and an invalid URL.

### *mod5.exe*
Executable file of cross-compiled Go code for Mac/Windows.

