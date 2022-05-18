import logging
from slack_bolt import App
from botway import GetToken, GetSigningSecret

logging.basicConfig(level=logging.DEBUG)

app = App(token=GetToken(), signing_secret=GetSigningSecret())

@app.middleware
def log_request(logger, body, next):
    logger.debug(body)

    return next()

@app.command("/hello-bolt-python")
def hello_command(ack, body):
    user_id = body["user_id"]

    ack(f"Hi <@{user_id}>!")

@app.event("app_mention")
def event_test(body, say, logger):
    logger.info(body)
    say("What's up?")

@app.error
def global_error_handler(error, body, logger):
    logger.exception(error)
    logger.info(body)

if __name__ == "__main__":
    app.start(3000)
