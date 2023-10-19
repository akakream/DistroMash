import pandas as pd

data = """
HH:MM:SS   Kbps in  Kbps out
08:19:31      1.54      0.51
08:19:32      2.63     34.21
08:19:33      3.01      1.92
08:19:34     14.34     30.29
08:19:35     25.12     65.47
08:19:36     15.03     68.29
08:19:37     57.10     45.76
08:19:38      5.03      3.93
08:19:39      0.52      0.52
08:19:40     30.08     57.08
08:19:41      2.53      3.04
08:19:42     37.01     50.98
08:19:43      8.43      7.12
08:19:44      0.00      0.00
08:19:45    134.02  39828.93
08:19:46   1188.57  154134.9
08:19:47   1395.16  94711.09
08:19:48   1415.09  94706.24
08:19:49   1413.92  94704.99
08:19:50   1416.71  94708.02
08:19:51   1883.36  94886.62
08:19:52   1495.37  94856.34
08:19:53   1478.10  94792.70
08:19:54   1484.53  94804.48
08:19:55   1468.03  94787.97
08:19:56   1467.99  94811.54
08:19:57   1587.05  94787.17
08:19:58   1486.79  94804.68
08:19:59   1485.82  94803.21
08:20:00   1499.11  94821.77
08:20:01   1490.67  94814.32
08:20:02   1500.65  94807.23
08:20:03   1477.78  94795.39
08:20:04   2360.02  94876.77
08:20:05   1092.07  94430.51
08:20:06   1444.17  94694.22
08:20:07   1480.29  94766.75
08:20:08   1397.62  94751.06
08:20:09    796.45  94398.85
08:20:10    693.47  94415.65
08:20:11   1121.12  94401.03
08:20:12   1135.04  94482.56
08:20:13   1101.53  94418.76
08:20:14   1119.71  94413.70
08:20:15    977.28  94387.25
08:20:16    927.32  94241.69
08:20:17    962.67  94166.02
08:20:18   1034.68  94271.38
08:20:19   2354.68  94369.67
08:20:20   1144.36  94311.63
08:20:21    945.75  94368.59
08:20:22   1043.36  94288.40
08:20:23   1029.35  94383.17
08:20:24    785.94  94155.77
08:20:25    874.84  94143.10
08:20:26    948.42  94138.82
08:20:27   1036.83  94244.21
08:20:28   1128.45  94351.76
08:20:29   1174.07  94470.05
08:20:30   1101.81  94330.47
08:20:31   1125.48  94398.02
08:20:32    821.41  94135.73
08:20:33    883.69  94193.34
08:20:34    853.23  94104.22
08:20:35    931.41  94088.72
08:20:36   1034.29  94271.54
08:20:37    945.26  94245.55
08:20:38   1033.53  94216.28
08:20:39    878.06  94118.84
08:20:40    950.02  94254.75
08:20:41    514.67  93737.30
08:20:42    517.07  93753.64
08:20:43    518.34  93753.14
08:20:44    482.06  93632.78
08:20:45    505.27  93614.00
08:20:46    286.53  44385.99
08:20:47      1.33      1.33
08:20:48     22.11     31.72
08:20:49      0.52      0.52
08:20:50      2.19      4.96
08:20:51     19.61     54.68
08:20:52     12.96     13.74
08:20:53     40.72     51.94
08:20:54     14.21     51.86
08:20:55     10.60     42.57
08:20:56      8.92     12.76
08:20:57     28.25     65.24
08:20:58      3.01      1.92
08:20:59      0.00      0.00
08:21:00     11.65     12.33
08:21:01     40.68     67.46
08:21:02      6.20      7.58
08:21:03      2.50      1.40
08:21:04     11.64      3.12
08:21:05     28.95     54.03
08:21:06     38.26     58.82
08:21:07    128.16    142.37
08:21:08     50.35     92.34
08:21:09     51.18     92.14
08:21:10     20.80     23.37
08:21:11     11.37     43.62
08:21:12     39.58     31.88
08:21:13      3.01      1.92
08:21:14      1.54     17.93
08:21:15      6.42      7.57
08:21:16     15.82     46.58
08:21:17     13.97     45.24
08:21:18      3.01      2.98
08:21:19     11.64      3.13
08:21:20      3.78      2.67
08:21:21      4.47      6.36
08:21:22     18.44     67.81
08:21:23      8.95     10.96
08:21:24     22.31     39.46
08:21:25     23.41     78.46
08:21:26     65.32     79.88
08:21:27     35.92     58.41
08:21:28     21.08     13.79
08:21:29      0.00      0.00
08:21:30     12.04     16.90
08:21:31      6.96      6.75
08:21:32     27.42     75.79
08:21:33      2.50      1.40
08:21:34      0.51      0.51
08:21:35     14.06     17.64
08:21:36     22.10     45.96
08:21:37    126.15    103.93
08:21:38     14.61     42.82
08:21:39     22.10     40.42
08:21:40     20.47     56.68
08:21:41      0.51      2.09
08:21:42     29.55     64.05
08:21:43      8.71      8.62
08:21:44      0.00      0.00
08:21:45      9.00     26.70
08:21:46     48.40     67.78
08:21:47     11.86     11.55
08:21:48    668.79  147203.4
08:21:49   1158.62  94506.48
08:21:50   1161.93  94470.64
08:21:51   1267.56  94552.61
08:21:52   1165.89  94444.36
08:21:53   1146.39  94414.23
08:21:54   1133.15  94375.36
08:21:55   1140.52  94377.09
08:21:56   1136.05  94335.29
08:21:57   1181.06  94375.82
08:21:58   1302.12  94538.77
08:21:59   1514.58  94771.92
08:22:00   1575.25  94621.76
08:22:01   1467.43  94793.21
08:22:02   1514.53  94845.33
08:22:03   1491.47  94838.42
08:22:04   1852.55  94404.67
08:22:05   1450.67  94765.56
08:22:06   1484.79  94757.92
08:22:07   1491.85  94792.73
08:22:08   1497.52  94799.72
08:22:09   1547.75  94850.29
08:22:10   1594.65  95019.12
08:22:11   1630.25  94979.81
08:22:12   1591.18  95013.78
08:22:13    891.20  94406.48
08:22:14   1097.02  94726.80
08:22:15   1314.85  94759.48
08:22:16    849.10  94812.44
08:22:17   1210.23  94973.01
08:22:18   1522.08  94872.01
08:22:19   2413.36  94842.73
08:22:20   1495.03  94875.65
08:22:21   1452.51  94950.68
08:22:22    915.56  94450.06
08:22:23   1135.46  94635.41
08:22:24   1660.67  95004.39
08:22:25   1029.20  94820.13
08:22:26   1700.30  94979.30
08:22:27   1586.45  94994.33
08:22:28   1081.26  94860.80
08:22:29   1440.32  94839.32
08:22:30   1069.12  94468.15
08:22:31   1408.37  94692.57
08:22:32   1435.74  94769.71
08:22:33   1495.44  94818.71
08:22:34   1650.82  94999.25
08:22:35   1501.77  94866.73
08:22:36   1427.40  94809.05
08:22:37   1211.92  94865.81
08:22:38   1368.13  94871.40
08:22:39    843.64  94296.22
08:22:40   1002.82  94300.76
08:22:41   1007.39  94330.32
08:22:42   1006.14  94297.79
08:22:43    975.81  94337.55
08:22:44    972.84  94295.87
08:22:45    863.38  94342.42
08:22:46   1009.54  94304.64
08:22:47    686.93  94025.03
08:22:48    479.88  90164.45
08:22:49     17.56     18.43
08:22:50      8.42     40.53
08:22:51     10.95     45.39
08:22:52     23.08     15.77
08:22:53      4.55      4.16
08:22:54      0.52      0.52
08:22:55      7.18      6.42
08:22:56     67.01     80.37
08:22:57     52.45    121.80
08:22:58      2.50      1.41
08:22:59      6.81      6.91
08:23:00      8.34     40.05
08:23:01      3.66     36.24
08:23:02      0.00      0.00
08:23:03      9.22      7.63
08:23:04     23.50     14.54
08:23:05     30.82     40.99
08:23:06     30.49     45.09
08:23:07     41.06     68.07
08:23:08     21.08     14.89
08:23:09      7.54     11.81
08:23:10     40.80     84.60
08:23:11     46.43     81.71
08:23:12     47.01    113.10
08:23:13      8.43      7.19
08:23:14      0.51      0.51
08:23:15      1.25      0.70
08:23:16     17.65     19.86
08:23:17      8.84     40.57
08:23:18      8.43      7.12
08:23:19     28.18     21.25
08:23:20     14.97     13.59
08:23:21      2.11     33.50
08:23:22      6.33     53.27
08:23:23      3.01      5.24
08:23:24      2.78      1.97
08:23:25     24.05     35.10
08:23:26     64.79     53.42
08:23:27     35.92     86.87
08:23:28     51.91     77.86
08:23:29      7.52      7.91
08:23:30      5.20      6.65
08:23:31      3.14     33.50
08:23:32      1.33     33.71
08:23:33      2.50      1.40
08:23:34     24.91     19.22
08:23:35     23.06     25.51
08:23:36      0.94      0.94
08:23:37     66.46     53.23
08:23:38      8.12     37.41
08:23:39      7.00      6.88
08:23:40      3.26      3.23
08:23:41     45.04     83.18
08:23:42     19.68     37.20
08:23:43     18.34     87.86
08:23:44      1.03      2.09
08:23:45     22.18     32.75
08:23:46     19.97     15.39
08:23:47     29.84     20.77
08:23:48     22.70     47.54
08:23:49      5.93      5.77
08:23:50    112.01  41869.97
08:23:51   1085.17  149355.2
08:23:52   1253.28  94529.86
08:23:53   1114.41  94400.49
08:23:54   1115.47  94369.58
08:23:55    999.01  94260.55
08:23:56    908.59  94144.07
08:23:57   1008.24  94152.08
08:23:58    955.65  94244.96
08:23:59    909.17  94037.60
08:24:00    908.31  94121.00
08:24:01   1010.31  94271.69
08:24:02   1150.31  94393.47
08:24:03   1337.54  94644.23
08:24:04   2025.60  94833.86
08:24:05   1671.75  95000.28
08:24:06   1481.97  94921.31
08:24:07   1369.43  94802.80
08:24:08   1408.53  94796.97
08:24:09   1374.62  94821.41
08:24:10   1808.76  95056.10
08:24:11   1546.88  94867.21
08:24:12   1562.83  94823.66
08:24:13   1493.95  94780.80
08:24:14   1530.52  94858.16
08:24:15   1480.25  94810.92
08:24:16   1462.23  94777.00
08:24:17   1522.65  94843.62
08:24:18   1501.04  94793.51
08:24:19   2510.44  94820.11
08:24:20   1336.67  94682.15
08:24:21   1253.17  94559.48
08:24:22   1372.91  94719.75
08:24:23   1617.94  94912.59
08:24:24   1707.48  94995.99
08:24:25   1540.39  94851.98
08:24:26    735.89  94331.75
08:24:27    877.01  94333.37
08:24:28    891.69  94309.19
08:24:29    861.54  94347.51
08:24:30   1048.66  94375.04
08:24:31    974.68  94273.66
08:24:32   1088.67  94371.97
08:24:33    945.35  94201.08
08:24:34    985.17  94214.46
08:24:35   1068.70  94364.87
08:24:36    984.59  94213.70
08:24:37    959.63  94540.73
08:24:38   1056.48  94311.97
08:24:39   1261.23  94544.94
08:24:40   1023.08  94552.34
08:24:41    849.85  94239.27
08:24:42    864.51  94233.10
08:24:43    895.28  94148.50
08:24:44    713.62  93986.06
08:24:45    507.04  93788.53
08:24:46    408.48  93775.29
08:24:47    491.42  93767.82
08:24:48    487.03  93781.91
08:24:49    505.39  93875.90
08:24:50    477.21  93757.55
08:24:51    288.23  49910.29
08:24:52     18.30     46.97
08:24:53      7.68      8.31
08:24:54     18.97     13.50
08:24:55     17.33    259.75
08:24:56     29.62     48.98
08:24:57      7.47     11.89
08:24:58     35.34     55.89
08:24:59     29.95     72.66
08:25:00      5.72      7.58
08:25:01     13.46     17.54
08:25:02      5.36     35.37
08:25:03     21.46     14.39
08:25:04     17.60      9.99
"""

# Split the data into lines
lines = data.strip().split("\n")
header = lines[1].split()

# Create an empty list to store the data
data_list = []

# Start from line 2 to extract the data
for line in lines[2:]:
    parts = line.split()
    data_dict = dict(zip(header, parts))
    data_list.append(data_dict)

# Create a DataFrame from the list of dictionaries
df = pd.DataFrame(data_list)

# Save the DataFrame to a CSV file
df.to_csv("networkDataDockerRegistry.csv", index=False)
