import ctypes

# Завантажте вашу спільну бібліотеку
find_free_port_lib = ctypes.CDLL('./find_free_port.dll')

# Налаштування типів аргументів і результатів для виклику функції
find_free_port_lib.FindFreePort.restype = ctypes.c_int

# Виклик функції
free_port = find_free_port_lib.FindFreePort()
print(f'Free port: {free_port}')