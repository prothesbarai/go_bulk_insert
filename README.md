# 📦 Go Bulk Insert

🚀 Golang দিয়ে দ্রুত এবং efficient ভাবে **Bulk Insert (multiple data insert)** করার একটি simple utility।

---

## ✨ Features

- ⚡ Fast bulk insert
- 🧠 Simple & clean code structure
- 🔄 Multiple row একসাথে insert
- 💾 Database performance optimized

---

## 🛠️ Requirements

- Go (Golang) 1.18+
- SQL Database (MySQL / PostgreSQL / MSSQL)

---

## 📂 Project Structure

```
go_bulk_insert/
│── main.go
│── bulk_insert.go
│── db.go
│── README.md
```

---

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/prothesbarai/go_bulk_insert.git
cd go_bulk_insert
```

---

### 2️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 3️⃣ Database Config Setup

`db.go` ফাইলে তোমার database credentials বসাও:

```go
dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
```

---

## 📌 Usage Example

```go
rows := [][]interface{}{
    {"John", 25},
    {"Alice", 30},
    {"Bob", 28},
}

err := BulkInsert(db, "users", []string{"name", "age"}, rows)
if err != nil {
    log.Fatal(err)
}
```

---

## ⚙️ How It Works

- একাধিক row collect করা হয়  
- SQL query dynamically generate করা হয়  
- একবারে সব data insert করা হয়  

---

# Flutter UI Side 

- Install Package in pubspec.yaml File

```yaml
  file_picker: ^10.3.10
  csv: ^6.0.0
```
- bulk_insert_screen.dart (page code)

```dart
  File? selectedFile;
  List<Map<String,dynamic>> parseProductListData = [];
  bool isLoading = false;
  static const int maxProducts = 50000;

  /// >>> Pick CSV file ========================================================
  Future<void> pickCSVFile() async{
    final result = await FilePicker.platform.pickFiles(type: FileType.custom,allowedExtensions: ['csv','json']);
    if(result != null){
      selectedFile = File(result.files.single.path!);
      // >>> File size check ===================================================
      int fileSizeInBytes = selectedFile!.lengthSync();
      double fileSizeInMB = fileSizeInBytes / (1024 * 1024); // >>> Convert Bytes to MB
      if(!mounted) return;
      if(fileSizeInMB > 70){
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("File size is too large. Maximum allowed is 70 MB")));
        return; //>>> stop processing
      }
      // <<< File size check ===================================================

      if(!mounted) return;
      setState(() { isLoading = true; });

      // >>> For CSV Type File Purpose =========================================
      if(selectedFile!.path.endsWith('.csv')){
        final input = selectedFile!.openRead();
        final fields = await input.transform(utf8.decoder).transform(const CsvToListConverter()).toList();
        parseProductListData = [];
        for(int i = 1; i < fields.length && parseProductListData.length < maxProducts; i++){
          parseProductListData.add({
            "name": fields[i][0].toString(),
            "price": double.tryParse(fields[i][1].toString()) ?? 0.0,
          });
        }
        // >>> Limit Alert ui
        if(!mounted) return;
        if(parseProductListData.length == maxProducts){
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("Maximum $maxProducts products allowed per file")));
        }
      }
      // <<< For CSV Type File Purpose =========================================

      // >>> For JSON Type File Purpose ========================================
      else if(selectedFile!.path.endsWith('.json')){
        final response = ReceivePort();
        await Isolate.spawn(_jsonIsolate, [response.sendPort, selectedFile!.path]);
        parseProductListData = await response.first as List<Map<String,dynamic>>;
      }
      // <<< For JSON Type File Purpose ========================================
      if(!mounted) return;
      setState(() {isLoading = false;}); // >>> For UI Change Purpose
    }
  }
  // >>> For JSON Type File Purpose ============================================
  static void _jsonIsolate(List<dynamic> args) async {
    SendPort sendPort = args[0];
    String path = args[1];
    File file = File(path);
    String content = await file.readAsString();
    List<dynamic> jsonData = jsonDecode(content);
    List<Map<String,dynamic>> result = [];

    for(var item in jsonData){
      if(result.length >= maxProducts) break; //>>> limit enforce
      result.add({
        "name": item['name'].toString(),
        "price": double.tryParse(item['price'].toString()) ?? 0.0,
        // "col3": item['col3'].toString(),
        // ...............................
        // "col10": item['col10'].toString()
      });
    }
    sendPort.send(result);
  }
  // <<< For JSON Type File Purpose ============================================
  /// <<< Pick CSV file ========================================================

  /// >>> Upload Product With Chunks ===========================================
  Future<void> uploadProductsInChunks() async{
    if (parseProductListData.isEmpty) return;
    setState(() {isLoading = true;});
    const int chunkSize = 500;
    for(int i=0; i<parseProductListData.length; i+=chunkSize){
      final chunk = parseProductListData.sublist(i, i + chunkSize > parseProductListData.length ? parseProductListData.length : i+chunkSize);
      await sendChunk(chunk); // >>> API Call Here
    }
    setState(() {isLoading = false;});
    if(!mounted) return;
    ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("All products uploaded successfully")),);
  }
  Future<void> sendChunk(List<Map<String,dynamic>> chunk) async{
    try{
      final request = await HttpClient().postUrl(Uri.parse("http://192.168.100.16:8080/products/bulk"));
      request.headers.set("Content-Type", "application/json");
      request.write(jsonEncode({"products": chunk,}));
      final response = await request.close();
      if (response.statusCode != 200) {
        throw Exception("Failed to upload chunk");
      }
    }catch(e){
      debugPrint("Chunk upload error: $e");
    }
  }
  /// <<< Upload Product With Chunks ===========================================
```
- Button onpress call & Pick File in Scaffold Widget
```dart
GestureDetector(
    onTap: ()=> pickCSVFile(),
    child: Container(
        width: double.infinity,
        padding: EdgeInsets.all(20.w),
        decoration: BoxDecoration(border: Border.all(color: AppColors.primaryColor), borderRadius: BorderRadius.circular(12.r),),
        child: Column(
            children: [
                Icon(Icons.upload_file, size: 40.sp,color: AppColors.primaryColor,),
                SizedBox(height: 10.h),
                Text(selectedFile == null ? "Upload CSV/JSON File" : selectedFile!.path.split('/').last,style: TextStyle(color: AppColors.primaryColor),),
            ],
        ),
    ),
),

ElevatedButton(
    onPressed:() {
        uploadProductsInChunks();
    },
    child: Text("Upload"),
),
```

---

## ✅ Solution (Correct URI)
- 🟢 IF Use Android Emulator
```dart
Uri.parse("http://10.0.2.2:8080/products/bulk")
```

- 🟢 IF iOS Simulator
```dart
Uri.parse("http://127.0.0.1:8080/products/bulk")
```

- 🟢 IF Use Real Device Android (Mobile)
- - 👉 PC এর IP দিতে হবে
```dart
Uri.parse("http://192.168.0.105:8080/products/bulk")  
```

---

## 📊 Why Bulk Insert?

| Normal Insert | Bulk Insert |
|--------------|------------|
| প্রতি row আলাদা query | একবারে সব insert |
| Slow | Fast ⚡ |
| বেশি DB load | কম DB load |

---

## ❗ Important Notes

- বড় data insert করলে batch ব্যবহার করো  
- SQL injection avoid করতে parameterized query ব্যবহার করো  
- Transaction ব্যবহার করলে আরো safe হবে  

---


## 🧑‍💻 Author

**Prothes Barai**  
Software Engineer 🚀  

---

## 📄 License

MIT License
