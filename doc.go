// Package recordio implements a file format for a sequence of variable-length data.
//
// The layout of each record is
//
//  +-----------------------------------------------------------|
//  | 4-bytes Magic number (including compression mode)         |
//  +-----------------------------------------------------------|
//  | 4-byte Uncompressed record size (uint32 in little-endian) |
//  +-----------------------------------------------------------|
//  | 4-byte Compressed record size (uint32 in little-endian)   |
//  +-----------------------------------------------------------|
//  | Record data (maximum 1024MB)                              |
//  +-----------------------------------------------------------|
//
// Example:
//  // Writing records.
//  f, _ := os.Create("data.rec")
//  defer f.Close()
//
//  w, _ := recordio.NewWriter(f, recordio.Snappy)
//  w.Write([]byte("this is a record"))
//  w.Write([]byte("this is a second record"))
//  ...
//
//  // Reading records.
//  f, _ := os.Open("data.rec")
//  defer f.Close()
//
//  r, _ := recordio.NewReader(f)
//  for {
//    rec, err := r.Read()
//    if err == io.EOF {
//      break
//    }
//    ...
//  }
package recordio
