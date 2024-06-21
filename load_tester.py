import asyncio
import websockets
import random
import time
from multiprocessing import Process

async def test_connection(uri, message, index):
    async with websockets.connect(uri) as websocket:
        total_elapsed_time = 0
        
        init = time.time() * 1000  # Start time in milliseconds
        

        for i in range(200):
            await asyncio.sleep(random.uniform(1.0, 3.0))
            
            if not websocket.open:
                print(f"Connection {index}: WebSocket is not open.")
                return

            start_time = time.time() * 1000  # Start time in milliseconds
            await websocket.send(message)
            response = await websocket.recv()
            
            end_time = time.time() * 1000  # End time in milliseconds

            elapsed_time = end_time - start_time
            total_elapsed_time += elapsed_time
            print(f"Connection {index}: Received in {elapsed_time:.2f} ms")
        
    
    average_elapsed_time = total_elapsed_time / 10
    print(f"Connection {index}: Average elapsed time: {average_elapsed_time:.2f} ms")
    
    end = time.time() * 1000  # End time in milliseconds
    total_elapsed_time = end - init
    print(f"Connection {index}: Total elapsed time: {total_elapsed_time:.2f} ms")

def main():
    uri = "wss://hackathon-backend-asyof5iquq-an.a.run.app/ws"  # Replace with your WebSocket server URI
    message = '{"type": "tweet", "action": "get_newest", "data": {"index": 0}}'
    num_connections = 300  # Number of connections to spawn

    processes = []
    for i in range(num_connections):
        p = Process(target=asyncio.run, args=(test_connection(uri, message, i),))
        processes.append(p)
        p.start()
        time.sleep(0.01)  # Delay between each connection creation

    for p in processes:
        p.join()

if __name__ == "__main__":
    asyncio.run(main())