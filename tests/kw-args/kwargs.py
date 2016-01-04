from __future__ import print_function

def foo(*args, **kwargs):
    s = "args=%s kwds=%s" % (args,kwargs)
    return s

if __name__ == "__main__":
    print(foo())
    print(foo(a=3))
    kw=dict(a=3)
    print(foo(**kw))
