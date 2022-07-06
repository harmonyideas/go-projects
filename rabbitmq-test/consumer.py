#!/usr/bin/env python
import pika, sys, os, json, sqlite3


def main():

    try:
        sqlcon = sqlite3.connect("test.db")
        cur = sqlcon.cursor()

        # Create table
        cur.execute(
            """CREATE TABLE if not exists contactform
               (email text, subject text, message text)"""
        )
        sqlcon.commit()

    except sqlite3.Error as error:
        print("Error whiel connecting to sqlite", error)

    sqlcolumns = ["Email", "Subject", "Message"]
    credentials = pika.PlainCredentials("test", "test")
    connection = pika.BlockingConnection(
        pika.ConnectionParameters(
            host="localhost", port="5672", virtual_host="/", credentials=credentials
        )
    )
    channel = connection.channel()
    channel.queue_declare(queue="TestQueue")

    def worker(body):
        contactform_dict = json.loads(body)
        data = tuple(contactform_dict[c] for c in sqlcolumns)
        try:
            cur.execute("INSERT INTO contactform VALUES (?,?,?)", data)
            sqlcon.commit()
            print(" [x] Received %r" % contactform_dict)

        except sqlite3.Error as error:
            print("Error executing INSERT statement")

    def callback(ch, method, properties, body):
        worker(body)

    channel.basic_consume(
        queue="TestQueue", on_message_callback=callback, auto_ack=True
    )

    print(" [*] Waiting for messages. To exit press CTRL C")
    channel.start_consuming()


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("Interrupted")
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
