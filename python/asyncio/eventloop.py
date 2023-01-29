import asyncio
import datetime

def display_date(end_time):
    print(datetime.datetime.now())
    loop = asyncio.get_event_loop()
    if (loop.time() + 1.0) < end_time:
        loop.call_later(1, display_date, end_time)
    else:
        loop.stop()

loop = asyncio.get_event_loop()
end_time = loop.time() + 5.0

loop.call_soon(display_date, end_time)

try:
    loop.run_forever()
finally:
    loop.close()