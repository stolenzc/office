import Quartz

d = Quartz.CGSessionCopyCurrentDictionary()
if d.get("CGSSessionScreenIsLocked", 0) == 0:
    print(False)
else:
    print(True)
