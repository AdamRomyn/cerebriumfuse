```plaintext
    __   ___ ____    ___ ____  ____  ____ __ __ ___ ___      _____ __ __  _____  ___
   /  ] /  _]    \  /  _]    \|    \|    |  |  |   |   |    |     |  |  |/ ___/ /  _]
  /  / /  [_|  D  )/  [_|  o  )  D  )|  ||  |  | _   _ |    |   __|  |  (   \_ /  [_
 /  / |    _]    /|    _]     |    / |  ||  |  |  \_/  |    |  |_ |  |  |\__  |    _]
/   \_|   [_|    \|   [_|  O  |    \ |  ||  :  |   |   |    |   _]|  :  |/  \ |   [_
\     |     |  .  \     |     |  .  \|  ||     |   |   |    |  |  |     |\    |     |
 \____|_____|__|\_|_____|_____|__|\_|____|\__,_|___|___|    |__|   \__,_| \___|_____|

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
```

## Getting Started

1. Clone the project repository:
   ```bash
   git clone https://github.com/AdamRomyn/cerebriumfuse.git
   ```
2. Navigate to the project directory:
   ```bash
   cd cerebriumfuse
   ```
3. Run the setup command to initialize the project cache folders and mount point (install make first if not installed):
   ```bash
   make setup
   ```
4. After setup completes, start the application:
   ```bash
   make run
   ```
5. If you run the project more than once, make sure to unmount before starting again to ensure proper functioning:
   ```bash
   make unmount
   ```

## Testing the Project

To test the project, just go through the folder /mnt/all-projects and run commands. You should get normal outputs.

To test the cache run the test command:

```bash
make test-cache
```

During testing, the script executes two reads on files of the same type in two distinct projects. Initially, the output indicates retrieval from the NFS (Network File System). Subsequently, upon the second read, the output reflects retrieval from the cache, signifying successful caching.

The program confines the cache size to two files. This setup becomes critical as two additional files are read, causing the initial Python file to be evicted from the cache. A successful test execution is validated by observing "Cache 2" and "Cache 1" at the end of the output. This outcome confirms the Python file's eviction from the cache after the subsequent reads.

To verify the effectiveness of the test, the final command concatenates all files in the cache directory. The absence of Python output in the last two lines confirms the functionality of caching.

## Additional Information

- **File Hashing:** Files are hashed to identify uniqueness when stored/pulled from the cache. Note that I sampled the file data to hash as it would be expensive to use the entire file data to create the hash.
- **RLU Caching:** The application utilizes the RLU (Recently Least Used) caching method for optimizing performance.
- **Go Newbie:** Note that this is my first go project. I am not 100% sure of all the intricacies of the language and it would not represent my knowledge. Just assume over time I would learn best practices and other parts of the language that I have not explored yet.
- **Print statements:** I left in print statements so that you can see what the fuse file system is doing. Normally I would not include these or would have just created an env var for debugging and if true only showed the debugging output.

## Additional Improvements

- While the current implementation is efficient, there's room for improvement, especially in the reading of which item to remove from the cache. Check the `cache_service` for potential enhancements, indicated by the `@todo` marker on line 74.
- I could have also added caching for the directory lookup. (Probably would not be that useful though)
- Maybe we could store files in memory. Could cap the amount to store to be very small and reserved for files either that are accessed within a 5 minute period or specific files specified in the program parameters.
- Depending on how big the cache file count or size limit is. Run a background task that runs every few minutes to cleanup the cache on files that have not been accessed within a long period. This will just reduce the amount of space that the system will use when not actively fetching and using files.
