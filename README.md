# Boring Project Litepage

Site is hosted via Github pages here: https://man-on-box.github.io/boring-project-litepage/

This is an experiment in building a simple site, with minimal dependencies.

## Lines of code

We can do a simple count of lines of code with this command, omitting some files that shouldn't be included in the count:

```bash
git ls-files | grep -Ev '\.(json|md|svg|yaml|gitignore|toml|mod|sum)$' | xargs wc -l
```
