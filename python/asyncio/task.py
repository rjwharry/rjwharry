import asyncio
import time


async def func():
    print("hello")
    await asyncio.sleep(1)


async def func2():
    print("world")
    await asyncio.sleep(2)


async def main():
    task1 = asyncio.create_task(func())
    task2 = asyncio.create_task(func2())

    await task1
    await task2


if __name__ == "__main__":
    start = time.time()
    asyncio.run(main())
    end = time.time()
    print(str(round(end - start)) + " sec")
