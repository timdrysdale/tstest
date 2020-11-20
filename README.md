# tstest
utility for recording websocket TS stream from crossbar to file, and scripts for load test 


## Usage

```
./tstest -path /out/pend14/video -file pend14.ts
```

## Analysing the results

```
tsanalyze pend14.ts
```

```
===============================================================================
|  TRANSPORT STREAM ANALYSIS REPORT                                           |
|=============================================================================|
|  Transport Stream Id: .......... 1 (0x0001)  |  Services: .............. 1  |
|  Bytes: ......................... 4,440,184  |  PID's: Total: .......... 4  |
|  TS packets: ....................... 23,618  |         Clear: .......... 4  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|-----------------------------------------------------------------------------|
|  Transport stream bitrate, based on ....... 188 bytes/pkt    204 bytes/pkt  |
|  User-specified: ................................... None             None  |
|  Estimated based on PCR's: ................ 1,110,731 b/s    1,205,262 b/s  |
|-----------------------------------------------------------------------------|
|  Broadcast time: ................................... 31 sec (0 min 31 sec)  |
|-----------------------------------------------------------------------------|
|  Srv Id  Service Name                              Access          Bitrate  |
|  0x0001  Service01 .................................... C    1,077,951 b/s  |
|                                                                             |
|  Note 1: C=Clear, S=Scrambled                                               |
|  Note 2: Unless specified otherwise, bitrates are based on 188 bytes/pkt    |
===============================================================================
```

The stream is error free.