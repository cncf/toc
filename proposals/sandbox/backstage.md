## Name of project:

Backstage

## Description (what it does, why it is valuable, its origin and history)

[Backstage](https://backstage.io) is an open platform for building developer portals. It’s based on the developer portal we’ve been using internally at Spotify for over four years. Backstage can be as simple as a services catalog or as powerful as the UX layer for your entire tech infrastructure.

It’s a centralized platform for integrating all your various services and tooling (CI/CD, monitoring, GCP, AWS, microservices, data pipelines, cloud cost, etc., etc., etc.) into a unified experience, enabling engineers to spin up and ship faster — and your organization to work at scale.

### Features

* Create and manage all of your organization’s software and microservices in one place
* Services catalog keeps track of all software and its ownership
* Visualizations provide information about your backend services and tooling, and help you monitor them
* A unified method for managing microservices offers both visibility and control 
* Preset templates allow engineers to quickly create microservices in a standardized way
* Centralized, full-featured technical documentation with integrated tooling that makes it easy for developers to set up, publish, and maintain alongside their code

### Benefits

* For engineering managers, it allows you to maintain standards and best practices across the organization, and can help you manage your whole tech ecosystem, from migrations to test certification.
* For end users (developers), it makes it fast and simple to build software components in a standardized way, and it provides a central place to manage all projects and documentation.
* For platform engineers, it enables extensibility and scalability by letting you easily integrate new tools and services (via plugins), as well as extending the functionality of existing ones. 
* For everyone, it’s a single, consistent experience that ties all your infrastructure tooling, resources, standards, owners, contributors, and administrators together in one place.

### Origins
Backstage was born out of necessity at Spotify. We found that as we grew, our infrastructure was becoming more fragmented, our engineers less productive. (“What are we using for microservices these days?” “What version of that framework is everyone on?” “Didn’t someone else already build a service like that?” “I can’t find documentation for anything!”)

We needed to bring order to a sprawling ecosystem, and we needed to make it easier for engineers to do their work without infrastructure tooling, context switching, and cognitive overload getting in their way. We looked for an external solution, but nothing fit the bill. 

So we decided to build our own. 

Once we saw how quickly Backstage was embraced by Spotify’s engineers — and how it helped us scale at speed, and safely — we knew it could do the same for other tech organizations.

We believe Backstage has the ability to become the trusted, standard UX layer across the broader infrastructure landscape. [Read more](https://engineering.atspotify.com/2020/03/17/what-the-heck-is-backstage-anyway/)

## Statement on alignment with CNCF charter mission

Backstage aims to provide a consistent UX layer that ties together functionality provided by other open source projects -- many of them from the CNCF landscape. Backstage provides an end-to-end extensible UX layer with touchpoints at each step in the App Delivery pipeline of developing, distributing, deploying, managing and operating secure cloud-native applications.

## Comparison with similar projects (inside or outside the CNCF), including what differentiates this project
Backstage was developed because there weren't any existing projects for managing micro services when Spotify got started. Before being released as open source, Backstage grew far beyond the original use-cases of managing and discovering micro services.

The next closest projects we were able to identify are GitLab (providing a flexible DevOps toolchain) and Linkerd (service mesh for Kubernetes). Backstage is different in that it is a platform for building developer portals where you can integrate tools needed for any software development workflow.

## Sponsor from TOC:
Chris Aniscyk 

## Preferred maturity level:
Sandbox

## License
Apache License 2.0

## Source control
http://github.com/spotify/backstage

## External dependencies (including licenses)

A full listing of external dependencies can be found in the [Backstage GitHub network dependency graph](https://github.com/spotify/backstage/network/dependencies)

For the sake of being transparent, we’ve included a subset of the dependencies listed there from package.json and /app/package.json.

| PACKAGE                                 | REPO                                          | LICENSE                                 |
|-----------------------------------------|-----------------------------------------------|-----------------------------------------|
| @hot-loader/react-dom                   | hot-loader/react-dom                          |                                         |
| @lerna/package-graph                    | lerna/lerna                                   | MIT License                             |
| @lerna/project                          | lerna/lerna                                   | MIT License                             |
| @material-ui/core                       | mui-org/material-ui                           | MIT License                             |
| @material-ui/icons                      | mui-org/material-ui                           | MIT License                             |
| @material-ui/lab                        | mui-org/material-ui                           | MIT License                             |
| @rollup/plugin-commonjs                 | rollup/plugins                                | MIT License                             |
| @rollup/plugin-json                     | rollup/plugins                                | MIT License                             |
| @rollup/plugin-node-resolve             | rollup/plugins                                | MIT License                             |
| @spotify/eslint-config                  | spotify/web-scripts                           | Apache License 2.0                      |
| @spotify/eslint-config-oss              | spotify/eslint-config-oss                     | Apache License 2.0                      |
| @spotify/prettier-config                | spotify/web-scripts                           | Apache License 2.0                      |
| @storybook/addon-actions                | storybookjs/storybook                         | MIT License                             |
| @storybook/addon-links                  | storybookjs/storybook                         | MIT License                             |
| @storybook/addon-storysource            | storybookjs/storybook                         | MIT License                             |
| @storybook/addons                       | storybookjs/storybook                         | MIT License                             |
| @storybook/react                        | storybookjs/storybook                         | MIT License                             |
| @sucrase/webpack-loader                 | alangpierce/sucrase                           | MIT License                             |
| @svgr/plugin-jsx                        | gregberge/svgr                                | MIT License                             |
| @svgr/plugin-svgo                       | gregberge/svgr                                | MIT License                             |
| @svgr/rollup                            | gregberge/svgr                                | MIT License                             |
| @svgr/webpack                           | gregberge/svgr                                | MIT License                             |
| @testing-library/cypress                | testing-library/cypress-testing-library       | MIT License                             |
| @testing-library/jest-dom               | testing-library/jest-dom                      | MIT License                             |
| @testing-library/react                  | testing-library/react-testing-library         | MIT License                             |
| @testing-library/react-hooks            | testing-library/react-hooks-testing-library   | MIT License                             |
| @testing-library/user-event             | testing-library/user-event                    | MIT License                             |
| @types/body-parser                      | PatrickJS/types-body-parser                   | MIT License                             |
| @types/classnames                       | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/codemirror                       | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/color                            | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/compression                      | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/cookie-parser                    | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/cors                             | typed-typings/npm-cors                        | MIT License                             |
| @types/d3-force                         |                                               |                                         |
| @types/diff                             | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/express                          | types/express                                 |                                         |
| @types/express-serve-static-core        | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/fs-extra                         | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/git-url-parse                    | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/google-protobuf                  | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/helmet                           | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/html-webpack-plugin              | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/http-errors                      | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/http-proxy                       | typed-contrib/node-http-proxy                 | MIT License                             |
| @types/inquirer                         | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/jest                             | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/jquery                           | magnitronus/types-jquery                      |                                         |
| @types/jwt-decode                       |                                               |                                         |
| @types/lodash                           | typed-typings/npm-lodash                      | MIT License                             |
| @types/mini-css-extract-plugin          | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/morgan                           | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/node                             | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/node-fetch                       | nozzlegear/node-fetch.d.ts                    | MIT License                             |
| @types/nodegit                          | weswigham/AtTypes-nodegit                     | MIT License                             |
| @types/ora                              | sindresorhus/ora                              | MIT License                             |
| @types/passport                         | typed-typings/npm-passport                    | MIT License                             |
| @types/passport-github2                 | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/passport-google-oauth20          | typed-typings/npm-passport-google-oauth20     | MIT License                             |
| @types/passport-saml                    | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/react                            | Clever/minimal-react-typings                  |                                         |
| @types/react-dev-utils                  | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/react-dom                        | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/react-helmet                     | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/react-sparklines                 | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/recursive-readdir                |                                               |                                         |
| @types/rollup-plugin-peer-deps-external | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/rollup-plugin-postcss            | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/start-server-webpack-plugin      | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/stoppable                        | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/supertest                        | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/tar                              | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/uuid                             | huhuanming/uuid-types                         |                                         |
| @types/webpack                          | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/webpack-dev-server               | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/webpack-env                      | pictalk/types-webpack-env                     | MIT License                             |
| @types/webpack-node-externals           | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/yaml                             | eemeli/yaml                                   | ISC License                             |
| @types/yup                              | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| @types/zen-observable                   | DefinitelyTyped/DefinitelyTyped               | Other                                   |
| bfj                                     | philbooth/bfj                                 | MIT License                             |
| body-parser                             | expressjs/body-parser                         | MIT License                             |
| chalk                                   | chalk/chalk                                   | MIT License                             |
| chokidar                                | paulmillr/chokidar                            | MIT License                             |
| classnames                              | JedWatson/classnames                          | MIT License                             |
| clsx                                    | lukeed/clsx                                   | MIT License                             |
| color                                   | Qix-/color                                    | MIT License                             |
| commander                               | tj/commander.js                               | MIT License                             |
| compression                             | expressjs/compression                         | MIT License                             |
| cookie-parser                           | expressjs/cookie-parser                       | MIT License                             |
| cors                                    | expressjs/cors                                | MIT License                             |
| cross-env                               | kentcdodds/cross-env                          | MIT License                             |
| css-loader                              | webpack-contrib/css-loader                    | MIT License                             |
| cypress                                 | cypress-io/cypress                            | MIT License                             |
| d3-force                                | d3/d3-force                                   | BSD 3-Clause "New" or "Revised" License |
| dashify                                 | jonschlinkert/dashify                         | MIT License                             |
| del                                     | sindresorhus/del                              | MIT License                             |
| diff                                    | kpdecker/jsdiff                               | Other                                   |
| dockerode                               | apocas/dockerode                              | Apache License 2.0                      |
| eslint                                  | eslint/eslint                                 | MIT License                             |
| eslint-plugin-cypress                   | cypress-io/eslint-plugin-cypress              | MIT License                             |
| eslint-plugin-import                    | benmosher/eslint-plugin-import                | MIT License                             |
| eslint-plugin-monorepo                  | azz/eslint-plugin-monorepo                    | MIT License                             |
| express                                 | expressjs/express                             | MIT License                             |
| express-promise-router                  | express-promise-router/express-promise-router | MIT License                             |
| fork-ts-checker-webpack-plugin          | TypeStrong/fork-ts-checker-webpack-plugin     | MIT License                             |
| fs-extra                                | jprichardson/node-fs-extra                    | MIT License                             |
| get-port                                | sindresorhus/get-port                         | MIT License                             |
| git-url-parse                           | IonicaBizau/git-url-parse                     | MIT License                             |
| globby                                  | sindresorhus/globby                           | MIT License                             |
| graphiql                                | graphql/graphiql                              | MIT License                             |
| graphql                                 | graphql/graphql-js                            | MIT License                             |
| handlebars                              | handlebars-lang/handlebars.js                 | MIT License                             |
| helmet                                  | helmetjs/helmet                               | MIT License                             |
| html-webpack-plugin                     | jantimon/html-webpack-plugin                  | MIT License                             |
| http-errors                             | jshttp/http-errors                            | MIT License                             |
| husky                                   | typicode/husky                                | MIT License                             |
| inquirer                                | SBoudrias/Inquirer.js                         | MIT License                             |
| jest                                    | facebook/jest                                 | MIT License                             |
| jest-css-modules                        | justinsisley/Jest-CSS-Modules                 |                                         |
| jest-esm-transformer                    | ActuallyACat/jest-esm-transformer             |                                         |
| jest-fetch-mock                         | jefflau/jest-fetch-mock                       | MIT License                             |
| jose                                    | panva/jose                                    | Other                                   |
| jwt-decode                              | auth0/jwt-decode                              | MIT License                             |
| knex                                    | knex/knex                                     | MIT License                             |
| lerna                                   | lerna/lerna                                   | MIT License                             |
| lint-staged                             | okonet/lint-staged                            | MIT License                             |
| lodash                                  | lodash/lodash                                 | Other                                   |
| material-table                          | mbrn/material-table                           | MIT License                             |
| mini-css-extract-plugin                 | webpack-contrib/mini-css-extract-plugin       | MIT License                             |
| moment                                  | moment/moment                                 | MIT License                             |
| morgan                                  | expressjs/morgan                              | MIT License                             |
| msw                                     | mswjs/msw                                     | MIT License                             |
| node-fetch                              | node-fetch/node-fetch                         | MIT License                             |
| nodegit                                 | nodegit/nodegit                               | MIT License                             |
| nodemon                                 | remy/nodemon                                  | MIT License                             |
| ora                                     | sindresorhus/ora                              | MIT License                             |
| passport                                | jaredhanson/passport                          | MIT License                             |
| passport-github2                        | ppremk/passport-github                        | MIT License                             |
| passport-google-oauth20                 | jaredhanson/passport-google-oauth2            | MIT License                             |
| passport-saml                           | bergie/passport-saml                          | MIT License                             |
| prettier                                | prettier/prettier                             | MIT License                             |
| prop-types                              | facebook/prop-types                           | MIT License                             |
| raw-loader                              | webpack-contrib/raw-loader                    | MIT License                             |
| rc-progress                             | react-component/progress                      | MIT License                             |
| react                                   | facebook/react                                | MIT License                             |
| react-dev-utils                         | facebook/create-react-app                     | MIT License                             |
| react-dom                               | facebook/react                                | MIT License                             |
| react-helmet                            | nfl/react-helmet                              | MIT License                             |
| react-hook-form                         | react-hook-form/react-hook-form               | MIT License                             |
| react-hot-loader                        | gaearon/react-hot-loader                      | MIT License                             |
| react-markdown                          | rexxars/react-markdown                        | MIT License                             |
| react-router                            | ReactTraining/react-router                    | MIT License                             |
| react-router-dom                        | ReactTraining/react-router                    | MIT License                             |
| react-sparklines                        | borisyankov/react-sparklines                  | MIT License                             |
| react-syntax-highlighter                | conorhastings/react-syntax-highlighter        | MIT License                             |
| react-test-renderer                     | facebook/react                                | MIT License                             |
| react-use                               | streamich/react-use                           | The Unlicense                           |
| recursive-readdir                       | jergason/recursive-readdir                    | MIT License                             |
| replace-in-file                         | adamreisnz/replace-in-file                    |                                         |
| rollup                                  | rollup/rollup                                 | Other                                   |
| rollup-plugin-dts                       | Swatinem/rollup-plugin-dts                    |                                         |
| rollup-plugin-esbuild                   | egoist/rollup-plugin-esbuild                  | MIT License                             |
| rollup-plugin-image-files               | bspaulding/rollup-plugin-image-files          | MIT License                             |
| rollup-plugin-peer-deps-external        | pmowrer/rollup-plugin-peer-deps-external      | MIT License                             |
| rollup-plugin-postcss                   | egoist/rollup-plugin-postcss                  | MIT License                             |
| rollup-plugin-typescript2               | ezolenko/rollup-plugin-typescript2            | MIT License                             |
| sqlite3                                 | mapbox/node-sqlite3                           | BSD 3-Clause "New" or "Revised" License |
| start-server-and-test                   | bahmutov/start-server-and-test                |                                         |
| start-server-webpack-plugin             | ericclemmons/start-server-webpack-plugin      | MIT License                             |
| stoppable                               | hunterloftis/stoppable                        | MIT License                             |
| storybook-dark-mode                     | hipstersmoothie/storybook-dark-mode           | MIT License                             |
| style-loader                            | webpack-contrib/style-loader                  | MIT License                             |
| sucrase                                 | alangpierce/sucrase                           | MIT License                             |
| supertest                               | visionmedia/supertest                         | MIT License                             |
| swr                                     | vercel/swr                                    | MIT License                             |
| tar                                     | npm/node-tar                                  | ISC License                             |
| ts-jest                                 | kulshekhar/ts-jest                            | MIT License                             |
| ts-loader                               | TypeStrong/ts-loader                          | MIT License                             |
| ts-node                                 | TypeStrong/ts-node                            | MIT License                             |
| typescript                              | microsoft/TypeScript                          | Apache License 2.0                      |
| url-loader                              | webpack-contrib/url-loader                    | MIT License                             |
| uuid                                    | uuidjs/uuid                                   | MIT License                             |
| webpack                                 | webpack/webpack                               | MIT License                             |
| webpack-dev-server                      | webpack/webpack-dev-server                    | MIT License                             |
| webpack-node-externals                  | liady/webpack-node-externals                  | MIT License                             |
| whatwg-fetch                            | github/fetch                                  | MIT License                             |
| winston                                 | winstonjs/winston                             | MIT License                             |
| yaml                                    | eemeli/yaml                                   | ISC License                             |
| yml-loader                              | nkt/yml-loader                                | MIT License                             |
| yn                                      | sindresorhus/yn                               | MIT License                             |
| yup                                     | jquense/yup                                   | MIT License                             |
| zen-observable                          | zenparsing/zen-observable                     | MIT License                             |
| zombie                                  | assaf/zombie                                  | MIT License                             |


## Initial committers
* Patrik Oldsberg (@Rugvip, Spotify)
* Fredrik Adelöw (@freben, Spotify)
* Raghunandan Balachandran (@soapraj, Spotify)
* Ben Lambert (@benjdlambert, Spotify)
* Marcus Eide (@marcuseide, Spotify)
* Niklas Ek (@nikek, Spotify)
* Stefan Ålund (@stefanalund, Spotify)
* Kat Zhou (@katz95, Spotify)

## Infrastructure requests (CI / CNCF Cluster)

None.

## Communication channels (slack, irc, mailing lists)
* Discord chat: https://discord.gg/EBHEGzX
* Newsletter: https://mailchi.mp/spotify/backstage-community

## Issue tracker (GitHub by default)
http://github.com/spotify/backstage/issues

## Website (current version will move to project.cncf.io, see here for guidelines)
http://backstage.io

## Release methodology and mechanics
Releases are done though GitHub Actions, where NPM packages are built and published to the Backstage namespace in NPM. The release process is documented at https://github.com/spotify/backstage/blob/master/docs/publishing.md

## Social media accounts
* https://twitter.com/spotifyeng

## Community size and any existing sponsorship
_(06, 24, 2020)_

- 4613 Stars
- 256 Forks
- 66 Contributors
- 13 Releases
- Plugins/Integrations:
  - [CircleCI](https://github.com/spotify/backstage/tree/master/plugins/circleci)
  - [Lighthouse](https://github.com/spotify/backstage/tree/master/plugins/lighthouse)
  - [Tech Radar](https://github.com/spotify/backstage/tree/master/plugins/tech-radar)
  - [Sentry](https://github.com/spotify/backstage/tree/master/plugins/sentry) 

## Adopters:
https://github.com/spotify/backstage/blob/master/ADOPTERS.md

## Project logo in svg format
![Backstage SVG logo](https://backstage.io/img/logo.svg)
