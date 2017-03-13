<img align="right" src="docs/logo.png">
# BackLun
## Universal backend server
### Description
Backlun — a set of various Back-end services (platforms) with a API, but missing the Front-end. You can create any Front-end using any framework, libraries, JS dialects.

#### Why?
Backlun will be useful for Front-end developers who are not familiar with Back-end development, but who want to create a Front-end for a real dynamic application. Backlun is suitable for developers who want to try out a new Front-end framework, but do not want to create a back-end part of the application for this. Backlun is a good choice for an employer who wants to test the knowledge of a hired Front-end developer.

#### Examples of using
You want to learn how to write a Front-end on a TypeScript or ELM. You read the React tutorial and want to write the first practical application. You want to hire a new developer, but he previously made only static pages and you want to test his ability to create an application that interacts with the server through API.

### Platforms
- **ToDo List** — Write down your goals and tasks, split them into categories and perform.
- **Blog** — Manage the blog, create new posts, add tags and comment on them.
- **Market** — Put the goods in the shopping cart, top up the account with virtual money and make purchases.
- **Forum** — Sign up, communicate at a forum, create topics and reply to them.
- **Blackjack** — Play the classic card game and try to get 21 points before your opponents.
- **Calendar** — Keep a calendar, write in it future and past events, both short and long.
- **Geopos** — Select a point on the map and find the distance from you to it.
- **OAuth** — Use the newest OAuth 2.0 authorization service.
- **Chat** — Use chat on WebSockets.

### Install
#### Downloading
Select the desired version on [release page](https://github.com/Vladimir37/Backlun/releases) or select it below:
- [Backlun Windows](https://github.com/Vladimir37)
- [Backlun Linux](https://github.com/Vladimir37)
- [Backlun OSX](https://github.com/Vladimir37)

Download the archive, unzip it to any convenient place

#### Running the app
- **Windows** — Execute the `Backlun.exe` inside the directory.
- **Linux and OSX** — Execute the `Backlun` through the terminal. Example command: `./Backlun`. If necessary, give the file the right to execute: `chmod +x Backlun`.

#### Running the platforms
After starting, you will see the console. Enter one of the following commands:
- **help** — Show full help
- **help &lt;platform&gt;** — Show help for the selected platform.
- **start &lt;platform&gt; [&lt;port&gt;]** — Run the selected platform. You can optionally specify a port.

Examples:
- `start blog` — Start blog platform on port 8000.
- `start forum 8080` — Start forum platform on port 8080.
- `help market` — Show help for market platform.

##### Note for the OAuth platform
Starting of the OAuth platform is different from other platforms.
- **start oauth [&lt;port&gt; &lt;host&gt; &lt;key&gt;]**

`key` - key file. The default file name is `key.json` and it should be located in the root folder `Backlun/`. Key file structure: 
```
{ 
    "cid": string, 
    "csecret": string
}
```

#### Compilation
You can not use binary files, but compile Backlun from the source code yourself. Download [Backlun source code](https://github.com/Vladimir37/Backlun/archive/master.zip). Unzip to any convenient place, go to the directory and run `go build Backlun.go`. If you do not have Golang installed, then you can install it according to the [instructions](https://golang.org/doc/install).

### Development
Developers: [Vladimir37](https://github.com/Vladimir37) (ToDo, Blog, Market, Forum, Blackjack, Calendar), [Jaime](https://github.com/asm-jaime) (Geopos, OAuth, Chat).

Graphics made with [Logo Maker](http://logomakr.com).