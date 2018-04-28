# TranslateBot
A simple SlackBot that translates English text to Japanese and any other language to English.

## How to deploy
#### Assumptions
You are building the bot an a Raspberry Pi or similar ARM device. If not, you must make minor changes to the Dockerfile.
1. `$ docker build -t translatebot .`
2. `$ docker run -d -e SLACK_TOKEN=xoxb-your-slack-token -e TRANSLATE_API_KEY=your-google-translate-api-key translatebot`
* Add `--restart always` to the docker run command to make the container reboot after restarts
