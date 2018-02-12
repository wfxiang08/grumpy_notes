from hello_b import get_key_value_b
global_meta = {}
def hello(name, metas):
    a = get_key_value_b(metas, name)
    print "hello, world: %s %s %s" % (name, metas, a)


def get_key_value(reqs, key):
    return "result: %s %s" % (reqs.get(key, ""), global_meta)
