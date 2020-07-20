# Prefer Discrete GPU

**This is a CGo package**

For laptops with multiple GPUs, its likely that applications won't launch using the discrete GPU on Windows, but rather the integrated GPU. This can lead to real-time applications like games under-performing by default, which is not the ideal user experience.

## Credits

* [Hakan Guleryuz](https://groups.google.com/forum/#!topic/golang-nuts/7OHZcXUegF0) for documenting how they were able to export the NVIDIA/AMD variables in Golang.
* [SFML Contributors](https://github.com/SFML/SFML/commit/9a453ed9e3846e9f7998295b8966428a9a0b86f6#diff-93134bfcdd8e19cbd5fe05a57a658950R63) for inspiring how this feature should be consumed (ie. opt-in imported package)
