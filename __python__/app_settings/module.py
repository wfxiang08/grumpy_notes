from hello import hello, get_key_value

d = {"a": 1, "b": "hh"}

hello("wangfei", d)

b = get_key_value(d, "b")
print b

dd = get_key_value(d, "d")
print dd
