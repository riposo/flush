# Flush Plugin

Flush endpoint plugin for [Riposo](https://github.com/riposo/riposo).

This core plugin adds a `POST /__flush__` endpoint which can be used to clear
stores as described in the original [Kinto](https://kinto.readthedocs.io/en/latest/api/1.x/flush.html)
documentation.

**WARNING**: this feature can be useful for development and testing, but should
not be enabled in production environments.
