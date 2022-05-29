import json


class Human:
    def __init__(self, name: str):
        self.name = name

    def say(self, msg: str):
        print("(" + self.name + ") " + msg)

    def my_profile(self):
        print(json.dumps({"name": self.name}))


def main():
    john = Human("john")
    john.say("hello")
    john.my_profile()


if __name__ == "__main__":
    main()
