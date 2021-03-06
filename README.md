# Flush Plugin

Flush endpoint plugin for [Riposo](https://github.com/riposo/riposo).

This core plugin adds a `POST /__flush__` endpoint which can be used to clear
stores as described in the original [Kinto](https://kinto.readthedocs.io/en/latest/api/1.x/flush.html)
documentation.

**WARNING**: this feature can be useful for development and testing, but should
not be enabled in production environments.

## License

Copyright 2021 Black Square Media Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this material except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
