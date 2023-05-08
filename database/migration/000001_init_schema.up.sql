CREATE TABLE `anggota` (
  `id_anggota` varchar(26) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `sex` varchar(255) NOT NULL,
  `telp` varchar(255) NOT NULL,
  `alamat` text NOT NULL,
  `email` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `bukus` (
  `id_buku` varchar(26) NOT NULL,
  `isbn` varchar(255) NOT NULL,
  `id_kategori` varchar(26) NOT NULL,
  `judul_buku` varchar(255) NOT NULL,
  `id_penulisbuku` varchar(26) DEFAULT NULL,
  `id_penerbitbuku` varchar(26) DEFAULT NULL,
  `tahun_terbit` varchar(255) DEFAULT NULL,
  `stok_buku` int(11) DEFAULT NULL,
  `rak_buku` varchar(255) DEFAULT NULL,
  `deskripsi_buku` text DEFAULT NULL,
  `gambar_buku` varchar(255) DEFAULT NULL,
  `kondisi_buku` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `dendas` (
  `id_denda` varchar(26) NOT NULL,
  `jumlah_denda` int(11) NOT NULL,
  `tglpinjam` datetime(3) DEFAULT NULL,
  `tglhrskembali` datetime(3) DEFAULT NULL,
  `tglkembali` datetime(3) DEFAULT NULL,
  `id_peminjaman` varchar(26) NOT NULL,
  `id_anggota` varchar(26) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `detail_pinjams` (
  `id_detailpinjam` varchar(26) NOT NULL,
  `id_buku` varchar(26) DEFAULT NULL,
  `id_peminjaman` varchar(26) DEFAULT NULL,
  `kondisi` varchar(255) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `jenis_bukus` (
  `id_jenis` varchar(26) NOT NULL,
  `jenis_buku` varchar(255) NOT NULL,
  `deskripsi` text DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `pegawais` (
  `id_pegawai` varchar(26) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(100) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `peminjamen` (
  `id_peminjaman` varchar(26) NOT NULL,
  `id_anggota` varchar(26) DEFAULT NULL,
  `tgl_pinjam` datetime(3) DEFAULT NULL,
  `tgl_hrs_kembali` datetime(3) DEFAULT NULL,
  `jaminan` varchar(255) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `penerbit_bukus` (
  `id_penerbit` varchar(26) NOT NULL,
  `penerbit_buku` varchar(255) NOT NULL,
  `alamat_penerbit` varchar(255) DEFAULT NULL,
  `telp_penerbit` varchar(255) DEFAULT NULL,
  `email_penerbit` varchar(255) DEFAULT NULL,
  `deskripsi` text DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `penulis_bukus` (
  `id_penulis` varchar(26) NOT NULL,
  `penulis_buku` varchar(255) NOT NULL,
  `alamat_penulis` varchar(255) DEFAULT NULL,
  `email_penulis` varchar(255) DEFAULT NULL,
  `deskripsi` text DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT current_timestamp(3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for table `anggota`
--
ALTER TABLE `anggota`
  ADD PRIMARY KEY (`id_anggota`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `bukus`
--
ALTER TABLE `bukus`
  ADD PRIMARY KEY (`id_buku`),
  ADD UNIQUE KEY `isbn` (`isbn`),
  ADD KEY `fk_bukus_penerbit_buku` (`id_penerbitbuku`),
  ADD KEY `fk_bukus_kategori_jenis` (`id_kategori`),
  ADD KEY `fk_bukus_penulis_buku` (`id_penulisbuku`);

--
-- Indexes for table `dendas`
--
ALTER TABLE `dendas`
  ADD PRIMARY KEY (`id_denda`),
  ADD KEY `fk_anggota_denda` (`id_anggota`),
  ADD KEY `fk_peminjaman_denda` (`id_peminjaman`);

--
-- Indexes for table `detail_pinjams`
--
ALTER TABLE `detail_pinjams`
  ADD PRIMARY KEY (`id_detailpinjam`),
  ADD KEY `fk_buku_detail_pinjam` (`id_buku`),
  ADD KEY `fk_peminjaman_detail_pinjam` (`id_peminjaman`);

--
-- Indexes for table `jenis_bukus`
--
ALTER TABLE `jenis_bukus`
  ADD PRIMARY KEY (`id_jenis`),
  ADD UNIQUE KEY `jenis_buku` (`jenis_buku`);

--
-- Indexes for table `pegawais`
--
ALTER TABLE `pegawais`
  ADD PRIMARY KEY (`id_pegawai`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `peminjamen`
--
ALTER TABLE `peminjamen`
  ADD PRIMARY KEY (`id_peminjaman`),
  ADD KEY `fk_anggota_meminjam` (`id_anggota`);

--
-- Indexes for table `penerbit_bukus`
--
ALTER TABLE `penerbit_bukus`
  ADD PRIMARY KEY (`id_penerbit`),
  ADD UNIQUE KEY `penerbit_buku` (`penerbit_buku`),
  ADD UNIQUE KEY `email_penerbit` (`email_penerbit`);

--
-- Indexes for table `penulis_bukus`
--
ALTER TABLE `penulis_bukus`
  ADD PRIMARY KEY (`id_penulis`),
  ADD UNIQUE KEY `penulis_buku` (`penulis_buku`),
  ADD UNIQUE KEY `email_penulis` (`email_penulis`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `bukus`
--
ALTER TABLE `bukus`
  ADD CONSTRAINT `fk_bukus_kategori_jenis` FOREIGN KEY (`id_kategori`) REFERENCES `jenis_bukus` (`id_jenis`),
  ADD CONSTRAINT `fk_bukus_penerbit_buku` FOREIGN KEY (`id_penerbitbuku`) REFERENCES `penerbit_bukus` (`id_penerbit`),
  ADD CONSTRAINT `fk_bukus_penulis_buku` FOREIGN KEY (`id_penulisbuku`) REFERENCES `penulis_bukus` (`id_penulis`);

--
-- Constraints for table `dendas`
--
ALTER TABLE `dendas`
  ADD CONSTRAINT `fk_anggota_denda` FOREIGN KEY (`id_anggota`) REFERENCES `anggota` (`id_anggota`),
  ADD CONSTRAINT `fk_peminjaman_denda` FOREIGN KEY (`id_peminjaman`) REFERENCES `peminjamen` (`id_peminjaman`);

--
-- Constraints for table `detail_pinjams`
--
ALTER TABLE `detail_pinjams`
  ADD CONSTRAINT `fk_buku_detail_pinjam` FOREIGN KEY (`id_buku`) REFERENCES `bukus` (`id_buku`),
  ADD CONSTRAINT `fk_peminjaman_detail_pinjam` FOREIGN KEY (`id_peminjaman`) REFERENCES `peminjamen` (`id_peminjaman`);

--
-- Constraints for table `peminjamen`
--
ALTER TABLE `peminjamen`
  ADD CONSTRAINT `fk_anggota_meminjam` FOREIGN KEY (`id_anggota`) REFERENCES `anggota` (`id_anggota`);
COMMIT;