# Dockerfile (misalnya, di ./app/Dockerfile atau di root proyek)

FROM python:3.13
# Atau versi Python lain yang Anda inginkan, misalnya 3.10-alpine

WORKDIR /app

# Salin file konfigurasi dependensi dan requirements.txt terlebih dahulu
# Ini memanfaatkan Docker cache untuk build yang lebih cepat jika dependensi tidak berubah
COPY requirements.txt .

# Instal uv di dalam kontainer
RUN pip install uv

# Gunakan uv untuk menginstal semua dependensi dari requirements.txt
# '--system' diperlukan karena kita menginstal ke lingkungan sistem dalam kontainer
RUN uv pip install -r requirements.txt --system

# Salin sisa kode aplikasi Anda ke dalam kontainer
COPY . .

# Paparkan port yang akan digunakan aplikasi Anda (jika aplikasi Anda adalah web server)
EXPOSE 80

# Perintah default untuk menjalankan aplikasi. Ini bisa di-override oleh docker-compose.yml
CMD ["python", "main.py"]