import json
import random
from datetime import datetime, timedelta

# Base Url
photo_urls = [
    "https://images.unsplash.com/photo-123",
    "https://images.unsplash.com/photo-456",
    "https://images.unsplash.com/photo-789",
    "https://images.pexels.com/photo-111",
    "https://images.pexels.com/photo-222"
]

# Category, SubCategory, Tag, Color, Sized list
categories = [[1, 2], [2, 3], [3, 4], [4, 5], [5, 6], [6, 7], [7, 8], [8, 9], [9, 10]]
subcategories = [[3, 4], [4, 5], [5, 6], [6, 7], [7, 8]]
sub_subcategories = [[5, 6], [6, 7], [7, 8], [8, 9]]
tags_list = ["trending", "new", "sale", "hot", "featured", "exclusive", "limited", "premium"]
colors_list = ["red", "blue", "green", "black", "white", "yellow", "purple", "orange"]
sizes_list = ["small", "medium", "large", "xl", "xxl"]
units = ["pcs", "kg", "g", "liter", "meter", "dozen"]
attributes_list = ["color,size", "size,material", "color,material", "weight,color"]

products = []
start_date = datetime(2026, 4, 8, 12, 0, 0)

for i in range(1, 100001):
    store_id = random.randint(101, 500)
    price = round(random.uniform(50, 5000), 2)
    purchase_price = round(price * random.uniform(0.5, 0.9), 2)
    discount = random.randint(0, 50)
    discount_type = random.choice(["%", "amount"])

    if discount_type == "%":
        discounted_price = round(price * (1 - discount / 100), 2)
    else:
        discounted_price = round(max(price - discount, purchase_price), 2)

    variant_product = random.choice([0, 1])

    # Random Photo Selected [1-5]
    num_photos = random.randint(1, 5)
    photos = random.sample(photo_urls, min(num_photos, len(photo_urls)))
    for p in range(len(photos)):
        photos[p] = photos[p] + f"_product_{i}.jpg"

    thumbnail = f"https://images.thumb.com/product_{i}_thumb.jpg"
    featured_img = f"https://images.featured.com/product_{i}_featured.jpg"

    # Random Tag
    num_tags = random.randint(2, 5)
    tags = ",".join(random.sample(tags_list, min(num_tags, len(tags_list))))

    # Color & Size
    num_colors = random.randint(1, 4)
    colors = ",".join(random.sample(colors_list, num_colors))

    num_sizes = random.randint(1, 3)
    variations = ",".join(random.sample(sizes_list, num_sizes))

    # Store Code
    store_code = f"ST{store_id}"

    # Category Id
    category = random.choice(categories)
    subcategory = random.choice(subcategories)
    sub_subcategory = random.choice(sub_subcategories)

    # Stock IN
    stock_in = random.randint(0, 500)

    # Featured Product (Almost 20%)
    featured = 1 if random.random() < 0.2 else 0

    # Published (Almost 95%)
    published = 1 if random.random() < 0.95 else 0

    # Trashed (Almost 2%)
    trashed = 1 if random.random() < 0.02 else 0

    # Date Time
    created_at = start_date - timedelta(days=random.randint(0, 365), hours=random.randint(0, 23))
    updated_at = created_at + timedelta(days=random.randint(0, 30))

    product = {
        "name": f"Product {i}",
        "store_id": store_id,
        "store_code": store_code,
        "category_id": str(category),
        "subcategory_id": str(subcategory),
        "sub_subcategory_id": str(sub_subcategory),
        "photos": photos,
        "thumbnail": thumbnail,
        "featured_img": featured_img,
        "video_link": f"https://video.store.com/product_{i}.mp4" if random.random() < 0.3 else "",
        "tags": tags,
        "description": f"This is product {i} description. A high quality product with best features.",
        "price": price,
        "purchase_price": purchase_price,
        "discount": discount,
        "discount_type": discount_type,
        "discounted_price": discounted_price,
        "sku": f"SKU{10000 + i}",
        "unit": random.choice(units),
        "weight": round(random.uniform(0.1, 10), 2),
        "variant_product": variant_product,
        "attributes": random.choice(attributes_list),
        "choice_options": variations,
        "colors": colors,
        "variations": variations,
        "published": published,
        "trashed": trashed,
        "stock_in": stock_in,
        "featured": featured,
        "created_by": random.randint(1, 10),
        "created_at": created_at.isoformat() + "Z",
        "updated_at": updated_at.isoformat() + "Z"
    }
    products.append(product)

# JSON Created
output = {"products": products}

with open("products_10000.json", "w", encoding="utf-8") as f:
    json.dump(output, f, indent=2, ensure_ascii=False)

print("✅ With 10,000 Product 'products_10000.json' Successfully Created!")
print(f"File Size : {len(json.dumps(output)) / (1024 * 1024):.2f} MB")