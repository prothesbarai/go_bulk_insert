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
  static const int maxProducts = 100000;

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
            "store_id": int.tryParse(fields[i][1].toString()) ?? 0,
            "store_code": fields[i][2].toString(),
            "category_id": fields[i][3].toString(),
            "subcategory_id": fields[i][4].toString(),
            "sub_subcategory_id": fields[i][5].toString(),
            "photos": fields[i][6].toString(),
            "thumbnail": fields[i][7].toString(),
            "featured_img": fields[i][8].toString(),
            "video_link": fields[i][9].toString(),
            "tags": fields[i][10].toString(),
            "description": fields[i][11].toString(),
            "price": double.tryParse(fields[i][12].toString()) ?? 0.0,
            "purchase_price": double.tryParse(fields[i][13].toString()) ?? 0.0,
            "discount": double.tryParse(fields[i][14].toString()) ?? 0.0,
            "discount_type": fields[i][15].toString(),
            "discounted_price": double.tryParse(fields[i][16].toString()) ?? 0.0,
            "sku": fields[i][17].toString(),
            "unit": fields[i][18].toString(),
            "weight": double.tryParse(fields[i][19].toString()) ?? 0.0,
            "variant_product": int.tryParse(fields[i][20].toString()) ?? 0,
            "attributes": fields[i][21].toString(),
            "choice_options": fields[i][22].toString(),
            "colors": fields[i][23].toString(),
            "variations": fields[i][24].toString(),
            "published": int.tryParse(fields[i][25].toString()) ?? 0,
            "trashed": int.tryParse(fields[i][26].toString()) ?? 0,
            "stock_in": int.tryParse(fields[i][27].toString()) ?? 0,
            "featured": int.tryParse(fields[i][28].toString()) ?? 0,
            "created_by": int.tryParse(fields[i][29].toString()) ?? 1,
            "created_at": fields[i][30].toString(),
            "updated_at": fields[i][31].toString(),
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
    final decoded = jsonDecode(content);
    List<dynamic> jsonData;
    // >>> Check if the decoded JSON is a Map with a "products" key
    if (decoded is Map<String, dynamic> && decoded['products'] is List) {
      jsonData = decoded['products'];
    } else if (decoded is List) {
      jsonData = decoded;
    } else {
      // Invalid format
      sendPort.send([]);
      return;
    }
    List<Map<String,dynamic>> result = [];

    for(var item in jsonData){
      if(result.length >= maxProducts) break; //>>> limit enforce
      result.add({
        "name": item['name']?.toString() ?? "",
        "store_id": int.tryParse(item['store_id'].toString()) ?? 0,
        "store_code": item['store_code']?.toString() ?? "",
        "category_id": item['category_id']?.toString() ?? "",
        "subcategory_id": item['subcategory_id']?.toString() ?? "",
        "sub_subcategory_id": item['sub_subcategory_id']?.toString() ?? "",
        "photos": item['photos']?.toString() ?? "",
        "thumbnail": item['thumbnail']?.toString() ?? "",
        "featured_img": item['featured_img']?.toString() ?? "",
        "video_link": item['video_link']?.toString() ?? "",
        "tags": item['tags']?.toString() ?? "",
        "description": item['description']?.toString() ?? "",
        "price": double.tryParse(item['price'].toString()) ?? 0.0,
        "purchase_price": double.tryParse(item['purchase_price'].toString()) ?? 0.0,
        "discount": double.tryParse(item['discount'].toString()) ?? 0.0,
        "discount_type": item['discount_type']?.toString() ?? "",
        "discounted_price": double.tryParse(item['discounted_price'].toString()) ?? 0.0,
        "sku": item['sku']?.toString() ?? "",
        "unit": item['unit']?.toString() ?? "",
        "weight": double.tryParse(item['weight'].toString()) ?? 0.0,
        "variant_product": int.tryParse(item['variant_product'].toString()) ?? 0,
        "attributes": item['attributes']?.toString() ?? "",
        "choice_options": item['choice_options']?.toString() ?? "",
        "colors": item['colors']?.toString() ?? "",
        "variations": item['variations']?.toString() ?? "",
        "published": int.tryParse(item['published'].toString()) ?? 0,
        "trashed": int.tryParse(item['trashed'].toString()) ?? 0,
        "stock_in": int.tryParse(item['stock_in'].toString()) ?? 0,
        "featured": int.tryParse(item['featured'].toString()) ?? 0,
        "created_by": int.tryParse(item['created_by']?.toString() ?? "1") ?? 1,
        "created_at": item['created_at']?.toString() ?? DateTime.now().toIso8601String(),
        "updated_at": item['updated_at']?.toString() ?? DateTime.now().toIso8601String(),
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
    // >>> For Show Success Dialogue From Response Golang ======================
    int totalInserted = 0;
    // <<< For Show Success Dialogue From Response Golang ======================
    for(int i=0; i<parseProductListData.length; i+=chunkSize){
      final chunk = parseProductListData.sublist(i, i + chunkSize > parseProductListData.length ? parseProductListData.length : i+chunkSize);
      // await sendChunk(chunk); // >>> API Call Here

      // >>> For Show Success Dialogue From Response Golang ====================
      final response = await sendChunk(chunk); // >>> API Call Here
      if (response != null) {
        totalInserted += (response['inserted'] as num).toInt();
      }else{
        if(!mounted) return;
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("Chunk upload failed!")));
        setState(() { isLoading = false; });
        return;
      }
      // <<< For Show Success Dialogue From Response Golang ====================
    }
    setState(() {isLoading = false;});
    if(!mounted) return;
    // >>> For Show Success Dialogue From Response Golang ======================
    showDialog(
      context: context,
      barrierDismissible: false,
      builder: (context) => Dialog(
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20.r)),
        child: Container(
          padding: EdgeInsets.all(20.w),
          decoration: BoxDecoration(color: Colors.white, borderRadius: BorderRadius.circular(20.r),),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Icon(Icons.check_circle_outline, color: AppColors.primaryColor, size: 60.sp),
              SizedBox(height: 15.h),
              Text("Bulk Upload Success!", style: TextStyle(color: AppColors.primaryColor, fontSize: 20.sp, fontWeight: FontWeight.bold),),
              SizedBox(height: 10.h),
              Text("$totalInserted products have been successfully uploaded.", style: TextStyle(color: AppColors.primaryColor.withValues(alpha: 0.9), fontSize: 14.sp), textAlign: TextAlign.center,),
              SizedBox(height: 20.h),
              ElevatedButton(
                style: ElevatedButton.styleFrom(backgroundColor: Colors.white, foregroundColor: AppColors.primaryColor, shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.r)),),
                onPressed: () {
                  setState(() {selectedFile = null;parseProductListData = [];});
                  Navigator.of(context).pop();
                },
                child: Padding(padding: EdgeInsets.symmetric(horizontal: 20.w, vertical: 10.h), child: Text("OK", style: TextStyle(fontWeight: FontWeight.bold)),),
              ),
            ],
          ),
        ),
      ),
    );
    // <<< For Show Success Dialogue From Response Golang ======================
  }
  Future<Map<String, dynamic>?> sendChunk(List<Map<String,dynamic>> chunk) async{
    try{
      final request = await HttpClient().postUrl(Uri.parse("http://192.168.100.16:8080/products/bulk"));
      request.headers.set("Content-Type", "application/json");
      request.write(jsonEncode({"products": chunk,}));
      final response = await request.close();

      // >>> Show UI From Go Success Response ==================================
      final responseBody = await response.transform(utf8.decoder).join();
      final Map<String, dynamic> res = jsonDecode(responseBody);
      // <<< Show UI From Go Success Response ==================================

      if (response.statusCode != 200) {
        throw Exception("Failed to upload chunk");
      }
      return res;
    }catch(e){
      debugPrint("Chunk upload error: $e");
      return null;
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
