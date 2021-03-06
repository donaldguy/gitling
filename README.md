# Gitling: **git** for **l**earn**ing**(s)

## Quick Start (Incredibly WIP)

```shell
git clone git@github.com:donaldguy/gitling.git
#install node if needed
npm install -g parcel-bundler
# install bazel per https://docs.bazel.build/versions/master/install.html
# install https://github.com/bazelbuild/bazel-watcher
# by cloning; bazel build //ibazel; cp bazel-bin/ibazel/$(uname -s | tr 'WDL' 'wdl')_amd64_pure_stripped/ibazel /usr/local/bin
make
```

And https://localhost:8080 should be serving an autorebuilding API server, which proxies to frontend dev server on http://localhost:1234

(if you put it in a GOPATH and `go get` or `dep ensure`, you can `go run ./server/server_example.go` instead; but where's the fun in that)

Note that right now if you make local changes to this README they won't be reflected in what the server loads until `git commit`ed to `HEAD`

## What?

Aspirationally: `gitling` is a bi-directional knowledge server for software teams. 

Mechanically: Gitling is a self-contained daemon which serves a realtime-editable ~wiki to and from the Markdown content of one or more git repos.

### Foundational Goals:

- Be the ✨best✨ tool for a team of people to read, write, and update ***non-autogenerated*** "software documentation": onboarding, design, process, runbooks, FAQs, etc.
- Never demand it be the only tool for any-of-the-above. Let people use their `less`s, `vim`s, `code`s, GitHub edit buttons or IDEs as befits their workflow
- Be a viable tool for on-the-fly capture for e.g. incident management, work journaling, meeting notes, etc.
- Support commit-to-trunk and commit-to-PR (for at least GitHub and VSTS, ideally also GitLab, Bitbucket, Phabricator, Gerrit, ReviewBoard, ... )

### Eventual/Maybe Goals

- Facilitate easy/rich cross-linking to Project Management, Code Review, etc. external systems (configurable/extensible automatic linkification)
- Facilitate rich embed of external else-formatted knowledge: videos, slide decks, tweets, stack overflow questions, etc.
- Provide facilities for navigating and finding what you need in a large corpus of such knowledge
- Provide facilities for asynchronous critique and improvement of documentation (e.g. "Google Doc-style comments", proposed changes)
- Real-time collaboration

### Non-goals:

- Be a Blog platform or CMS (Prefer: [https://www.netlifycms.org/](https://www.netlifycms.org/) )
- Be a full-fledged/full-lifecycle project management system (Prefer: [https://basecamp.com/](https://basecamp.com/); [https://www.notion.so/](https://www.notion.so/);  )

## Why? (and the name)

About 4 years ago, my worldview was impacted by Andrew Clay Shafer's (aka [@littleidea](https://twitter.com/littleidea)) 2013 Velocity talk "there is no talent shortage":

[![there is no talent shortage -- velocity NYC 2013 -- andrew clay shafer](https://img.youtube.com/vi/P_sWGl7MzhU/0.jpg)](https://www.youtube.com/watch?v=P_sWGl7MzhU)

It's a little meandering; I'd still encourage you to watch it if you have the time. My TL;DW is thus:

- "Culture", linguistically, is a system for cultivating (people, ideas, products)
- The market in "knowledge work" is a competition of cultures, largely in incentives: what people are enabled to and rewarded for focus(ing) on
- What enables a team in this market to outcompete (individuals and) other teams is the creation of a culture which is better than competitors at continuously (generating, capturing, sharing, and remembering) learning(s)
- (This is the same essential insight as/of "Lean" in both its Ohno manufacturing and Reis "Startup" incarnations. See also Senge, Marsick & Watkins, Womack, lean.org, Mike Rother, more)
- "You are either building a learning organization or you will be losing to someone who is"
- (It turns out that in addition to being a path to "winning", a collaborative Learning Culture is also a path to happy, fulfilled employees)

Along the way he touches on the idea of Marsick & Watkins's (DLOQ) dimension of "embedded system": "

> Necessary systems to share learning are created, maintained,
and integrated with work; employees have access to these
high- and low-technology systems

By no means is this (even close to) the most important of those dimensions, but I have a vague belief that it can help trojan-horse in the others. In short, it is my hypothesis that if you make it easy enough you can trick a product-development culture into viewing internal docs (secretly: the knowledge they contain) as another product or even *the* ***real*** *product* of a SaaS team. ("Documentation driven development")

Eventually you reach the 🌈🦄 of "The real [Your Company/Product Here] was the ~~friends~~ knowledge we  ~~made~~ learned along the way" (learning together is  also an okay way to make friends; This is *mostly* not an anti-capitalist plot). 

### The name

If you watch the talk above (or look at [the slides](https://www.slideshare.net/littleidea/there-is-no-talent-shortage-velocity-2013)) there is a recurrent image: a vintage photograph of the testing of an early [Gatling](https://en.wikipedia.org/wiki/Gatling_gun)-type cannon:

[![Early Gatling gun](https://farm1.staticflickr.com/118/270353459_a87b601981.jpg)](https://www.flickr.com/photos/13035641@N00/270353459)

The metaphor is a little violent for my taste, but it gets the point across: a game-changing tool for "competitive advantage".

I would aspire for this software to be one as well.

## How?

Inspiration:

- UX:
    + [https://typora.io/](https://typora.io/)
    + [https://www.notion.so/](https://www.notion.so/)

- Premise:

    + [https://ikiwiki.info/](https://ikiwiki.info/)
    + [https://hackmd.io/](https://hackmd.io/)
    + [https://wiki.js.org/](https://wiki.js.org/)
    + [https://github.com/timberio/gitdocs](https://github.com/timberio/gitdocs)

- Design:

    + [https://github.com/vjeantet/hugo-theme-docdock](https://github.com/vjeantet/hugo-theme-docdock)
    + [https://docs.microsoft.com/en-us/azure/](https://docs.microsoft.com/en-us/azure/networking/networking-overview)

Implementation:

- [https://github.com/laobubu/HyperMD](https://github.com/laobubu/HyperMD)
- [https://www.jstree.com/](https://www.jstree.com/)

maybe later
- [https://github.com/nathan-osman/StackTack](https://github.com/nathan-osman/StackTack)
- [https://isomorphic-git.org/](https://isomorphic-git.org/) (serverless version)
