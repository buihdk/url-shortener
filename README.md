# URL Shortener
1. **Submitted by: KHOA BUI**
2. **Time spent: 10 hours**

## How to run
```
  go build
  ./url-shortener -h
  ./url-shortener configure -a dogs -u http://www.dogs.com
  ./url-shortener run -p 8080
  open localhost:8080/dogs
```
## Requirements
The list of redirection should be maintained in a command line tool, what can:
- [x] Manipulate YAML config file. Where the redirection list peristently stored.
- [x] Implement append to the list: `./url-shortener configure -a dogs -u www.dogs.com` 
- [x] Implement remove from the list: `./url-shortener -d dogs`
- [x] List redirections: `./url-shortener -l`
- [x] Run HTTP server on a given port: `./url-shortener run -p 8080`
- [x] Prints usage info: `./url-shortener -h`

## Bonus
- [ ] Track number of times each redirection is used. When the user uses `urlshorten -l`, the user should see redirections ranked by how many times they have been used.
- [ ] Provide a default shortening, if no example is given. For example, if `dogs` is not given, generate a random alphanumeric string of length 8.
- [ ] Build a Handler that doesn't read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you.
