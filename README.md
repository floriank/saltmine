[![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/floriank/saltmine?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![Saltmine](assets/logo_256.png)
# Saltmine

This is Saltmine - an attempt to modularize Redmine and in particular the OpenProject fork of it. While being able to fork OpenProject again, I will not consider this as the technical goal should be to provide a modularized system which enforces a microservice architecture from the beginning.

Saltmine will be written in GoLang - it's frontend will be provided by a different repository.

## Private repo?

This is private until there is something presentable.

## Current goals

- [ ] introduce basic models for Project, Tickets and users

## Long term goals

- [ ] provide a fully open sourced modularized project collaboration tool based on services
- [ ] provides an easily deployable solution via containers
- strong integration:
    - [ ] Github/GitLab
    - [ ] Wunderlist
    - [ ] Evernote
    - [ ] BitBucket
    - [ ] CI systems
- [ ] provide a migration script from Redmine/OpenProject to Saltmine

## Even longer term goals

This is stuff I miss from Redmine and OP which would be neat if we can have that:

- [ ] Full [WAI](http://www.w3.org/WAI/) compliance from the start with the support of a component model in the frontend
- [ ] Fuzzy menu to quickly change to another part of the application
- [ ] Watchable objects (not only tickets)
- [ ] mentioning people - and using Markdown in tickets FFS
- [ ] project sharing between instances
- [ ] cross-project pull requests (from a single ticket, merge related PRs of multiple repositories at the same time)

### Notes on migration

Migration would be a killer feature to provide for users to switch to Saltmine. There are some points discussed as of yet:

- Migration would only be partially possible, as not all features are compatible
- Textile content would be converted to Markdown (`pandoc` can do this)

## Contributing

1. Fork it ( https://github.com/floriank/saltmine/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
