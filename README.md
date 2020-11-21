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

The stream is error free for this single case.


## Load test

the loadtest script draws 3*20 video feeds for a minute (>800MB of data total), then analyses them

Initially it appeared there were some errors to worry about ...

```|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 25  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 17  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 13  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 8  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 25  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 17  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 13  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 8  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 25  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 17  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: ................. 13  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 8  |         Scrambled: ...... 0  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 6  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 5  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 6  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 5  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 6  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 5  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 6  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 5  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 6  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 5  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 6  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 5  |         With PCR's: ..... 1  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |

```

tsresync reported that there were 12 synchronisation issues in the 60 streams

```
*** Synchronization lost after 17,116 TS packets
*** Got 0x04 instead of 0x47 at start of TS packet
*** Synchronization lost after 21,027 TS packets
*** Got 0x1E instead of 0x47 at start of TS packet
*** Synchronization lost after 12,239 TS packets
*** Got 0x4C instead of 0x47 at start of TS packet
*** Synchronization lost after 36,191 TS packets
*** Got 0xF5 instead of 0x47 at start of TS packet
*** Synchronization lost after 17,518 TS packets
*** Got 0x04 instead of 0x47 at start of TS packet
*** Synchronization lost after 20,836 TS packets
*** Got 0x1E instead of 0x47 at start of TS packet
*** Synchronization lost after 11,715 TS packets
*** Got 0x4C instead of 0x47 at start of TS packet
*** Synchronization lost after 36,191 TS packets
*** Got 0xF5 instead of 0x47 at start of TS packet
*** Synchronization lost after 16,942 TS packets
*** Got 0x04 instead of 0x47 at start of TS packet
*** Synchronization lost after 21,027 TS packets
*** Got 0x1E instead of 0x47 at start of TS packet
*** Synchronization lost after 12,024 TS packets
*** Got 0x4C instead of 0x47 at start of TS packet
*** Synchronization lost after 36,191 TS packets
*** Got 0xF5 instead of 0x47 at start of TS packet
```

After resynchronisation, there were no errors:

```
     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With invalid sync: .................. 0  |         Scrambled: ...... 0  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     With transport error: ............... 0  |         With PCR's: ..... 1  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |
|     Suspect and ignored: ................ 0  |         Unreferenced: ... 0  |

```

What causes this issue, and how badly does it affect the decoder?

One of the affected files, pend03a.ts was played back through VLC media player with no obvious visual degradation.

jsmpeg has a find-sync mode, so it should not be a fundamental issue.




