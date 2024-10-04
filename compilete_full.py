def convert_to_c_array(input_file, output_file):
    try:
        with open(input_file, 'rb') as f:
            byte_content = f.read()

        with open(output_file, 'w') as f_out:
            f_out.write('unsigned char index_html[] = {\n')
            for i, byte in enumerate(byte_content):
                if i % 12 == 0:
                    f_out.write('\n')
                f_out.write(f' 0x{byte:02x},')
            f_out.write('\n};\n')
            f_out.write(f'unsigned int index_html_len = {len(byte_content)};\n')

        print(f'File "{input_file}" successfully converted to C array and saved in "{output_file}".')

    except FileNotFoundError:
        print(f'Error: File "{input_file}" not found.')

if __name__ == '__main__':
    input_html = 'front/index.html'  # Шлях до HTML-файлу
    output_header = 'index_html.h'   # Шлях до файлу заголовку
    convert_to_c_array(input_html, output_header)
