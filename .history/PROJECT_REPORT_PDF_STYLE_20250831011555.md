---
title: "گزارش پروژه کارشناسی - سامانه مدیریت و جستجوی لاگ با Go"
author: "دانشجو"
date: "خرداد 1404"
lang: fa
fontsize: 12pt
---

# صفحه عنوان

- دانشکده فنی و مهندسی
- گروه مهندسی کامپیوتر و فناوری اطلاعات
- گزارش پروژه جهت اخذ درجه کارشناسی
- رشته مهندسی کامپیوتر (نرم افزار)

## عنوان

طراحی و پیاده‌سازی پنل ادمین سبک برای جستجوی لاگ‌ها (بدون نیاز به Elasticsearch)

## نگارش

نام دانشجو: ...

## استاد راهنما

نام استاد: ...

---

# تقدیم به

به پدر و مادر عزیزم و همه کسانی که با محبت و همراهی، مسیر رشد و یادگیری را برایم هموار کردند.

---

# سپاسگزاری

از استاد راهنمای گرامی بابت راهنمایی‌های ارزشمند، از دوستان و همکلاسی‌ها بابت همکاری‌ها، و از خانواده عزیزم بابت حمایت‌های بی‌دریغشان صمیمانه قدردانی می‌کنم.

---

# فهرست مطالب

- چکیده
- فصل اول: مبانی نظری و فناوری‌های مورد استفاده
  - زبان برنامه‌نویسی Go (Golang)
  - فریم‌ورک Echo (وب)
  - پنل ادمین GoAdmin و تم AdminLTE
  - پایگاه‌داده SQLite
  - HTML / CSS / JavaScript و کتابخانه‌های گراف (ChartJS, ECharts)
- فصل دوم: مستند پیاده‌سازی و تحلیل سیستم
  - معماری، ساختار پوشه‌ها و ماژول‌ها
  - جریان اجرا (Flow)
  - لایه Handler / Service / Search / Backscanner
  - طراحی API جستجو و پاسخ‌ها
  - پنل ادمین و صفحات کاربردی پروژه
- فصل سوم: ارزیابی، کارایی و تست‌ها
- فصل چهارم: نتیجه‌گیری و کارهای آینده
- منابع

---

# چکیده

در این پروژه، یک سامانه سبک برای جستجوی لاگ‌ها در مقیاس‌های کوچک تا متوسط طراحی و پیاده‌سازی شده است که نیاز به راه‌حل‌های سنگینی مانند Elasticsearch را برطرف می‌کند. سیستم با زبان Go توسعه یافته، از Echo برای وب‌سرور، از GoAdmin برای پنل ادمین و از SQLite به‌عنوان پایگاه‌داده‌ی سبک استفاده می‌کند. منطق جستجو با استفاده از اسکن معکوس فایل‌ها (backscanner) و پردازش موازی (goroutine) طراحی شده تا بدون مصرف زیاد حافظه، خطوط اخیر و مرتبط را سریع بیابد و بازگرداند. خروجی‌ها به‌شکل JSON و امکان دریافت اکسل فراهم است. این راهکار برای فروشگاه‌ها و تیم‌های کوچک که نیاز به مشاهده سریع لاگ‌ها دارند، مناسب است.

---

# فصل اول: مبانی نظری و فناوری‌های مورد استفاده

## زبان برنامه‌نویسی Go (Golang)

- تاریخچه و فلسفه: طراحی توسط گوگل با تأکید بر سادگی، کارایی و هم‌زمانی.
- دلایل انتخاب در این پروژه:
  - هم‌زمانی قدرتمند با goroutine و channel برای پردازش موازی فایل‌ها.
  - کارایی و مصرف حافظه پایین، مناسب برای IO-bound بودن مسئله.
  - دیپلوی آسان (باینری تک‌فایله) و اکوسیستم مناسب برای وب و پایگاه‌داده.
- نحو و ویژگی‌ها:
  - تایپ ایستا، مدیریت خطا ساده، ابزارهای استاندارد (go fmt, go test, go mod).

## فریم‌ورک Echo (وب)

- Echo یک فریم‌ورک وب سریع برای Go است؛ مسیرها، میان‌افزارها و هندلرها را ساده می‌کند.
- در این پروژه برای راه‌اندازی سرور HTTP و اتصال با GoAdmin به‌کار رفته است.

## پنل ادمین GoAdmin و تم AdminLTE

- GoAdmin پنل مدیریتی آماده بر بستر Go است؛ تم AdminLTE تجربه کاربری استانداردی ارائه می‌دهد.
- قابلیت‌هایی مانند تعریف صفحه، منو، جدول، فرم، و اعمال فیلترها را ساده می‌کند.

## پایگاه‌داده SQLite

- پایگاه‌داده سبکِ فایل‌محور برای محیط‌های کوچک؛ بدون نیاز به سرویس جداگانه.
- مدیریت schema از طریق migration های داخل پروژه.

## HTML / CSS / JavaScript و کتابخانه‌های گراف

- برای نمایش عناصر UI و نمودارها از ChartJS و ECharts استفاده شده است.

---

# فصل دوم: مستند پیاده‌سازی و تحلیل سیستم

## معماری کلان و ساختار پوشه‌ها

- لایه ارائه: GoAdmin + Echo (تعریف صفحات و ثبت مسیرها)
- لایه کاربرد: handler (تبدیل ورودی‌ها و واکشی سرویس)
- لایه دامنه: service و search (منطق جستجو)
- لایه داده: Tdata (راه‌اندازی DB و migration ها)

ساختار کلیدی پروژه:

```text
Tdata/           # DB و مهاجرت‌ها
Tpage/           # صفحات پنل و UI
Tmenu/           # منوهای پنل
handler/         # لایه هندلر/سرویس/موتور جستجو
  SearchLogsEng/
    backscanner/
    search/
    service/
    params/
main.go          # راه‌اندازی Echo و GoAdmin
logs/            # لاگ‌های نمونه
uploads/         # دارایی‌ها و آپلودها
```

## جریان اجرا (Flow)

1) کاربر در پنل GoAdmin به صفحه جستجو می‌رود یا درخواست POST به `/search` ارسال می‌کند.
2) `SearchLogHandlerExternal` ورودی فرم را به مدل `SearchLogRequest` تبدیل می‌کند.
3) سرویس `adminService` روش `GetFilesInFolder` را در `search` فراخوانی می‌کند.
4) موتور جستجو فایل‌ها را به‌صورت معکوس و موازی می‌خواند، فیلتر می‌کند و نتیجه را برمی‌گرداند.

## راه‌اندازی برنامه (برگزیده از main.go)

- پیکربندی GoAdmin با SQLite و تم `adminlte`، ثبت مسیرهای UI و API.
- ثبت مسیرها:
  - `/admin` (UI)
  - `/admin/healthCheck` (سلامت)
  - `/search` (جستجوی لاگ‌ها، POST)

## لایه Handler (adminHandler)

- تابع `SearchLogHandlerExternal` ورودی‌ها را می‌خواند: `year, month, day, logType, limit, searchKey, fileName, notInclude*`.
- حالت `countOnly` فقط شمارنده نتایج را برمی‌گرداند.
- پاسخ با `Content-Type: application/json` ارسال می‌شود.

## لایه Service (adminService)

- تعریف واسط `SearchLogs` و تزریق وابستگی.
- هماهنگ‌کننده بین Handler و منطق جستجو.

## موتور جستجو (search)

- `GetFileNames`: ساخت مسیر روزانه `logs/YYYY-MM-DD[/errors]`، فیلتر نام فایل‌ها، حذف فایل‌های خالی.
- `CreateReaderForFile`: ایجاد `backscanner.Scanner` برای هر فایل معتبر.
- `ProcessFile`: خواندن معکوس خطوط در goroutine ها، محدودسازی با semaphore.
- `CreateLine`: پارس خطوط JSON/متن و تبدیل `time` به نوع زمان.
- `ProcessLine2`: فیلتر on-the-fly با `SearchKey`/`NotIncludeSearchKey` و توقف در `limit`.
- مرتب‌سازی نهایی براساس زمان و تولید خروجی JSON خوانا.

## اسکن معکوس فایل (backscanner)

- خواندن chunk-محور از انتهای فایل، پیدا کردن `\n` و بازگردانی خطوط.
- کنترل حافظه با `MaxBufferSize`، مدیریت `io.EOF` و خط‌های بلند.

## API جستجو و نمونه پاسخ

- Endpoint: `POST /search`
- پارامترهای مهم: `year, month, day, logType, limit, searchKey, fileName, notInclude*`

نمونه درخواست:

```bash
curl -X POST http://localhost:1235/search \
  -d year=2024 -d month=06 -d day=10 \
  -d logType=errors -d limit=50 \
  -d searchKey=timeout
```

## پنل ادمین و صفحات کاربردی

- صفحات پیش‌فرض: Dashboard، Form، Table، User.
- صفحه جستجوی لاگ‌ها: ارسال فرم و نمایش نتایج.
- خروجی اکسل: دانلود فایل `Search_result.xlsx` پس از ایجاد.

---

# فصل سوم: ارزیابی، کارایی و تست‌ها

## کارایی

- IO-bound: گلوگاه دیسک؛ پردازش موازی و فیلتر زودهنگام throughput را بالا می‌برد.
- پارامترهای تیونینگ: ظرفیت کانال‌ها، تعداد goroutine ها، `DefultReadLine`.

## تست‌ها

- واحد: `GetFileNames`, `CreateLine`, فیلترها.
- ادغامی: ساخت پوشه‌های لاگ نمونه و فراخوانی `/search`.
- دستی: سناریوهای فایل خالی، فایل بزرگ، چندفایل همزمان.

---

# فصل چهارم: نتیجه‌گیری و کارهای آینده

- نتیجه: راهکاری سبک، ساده و مؤثر برای مشاهده و جستجوی لاگ بدون Elasticsearch.
- کارهای آینده: هایلایت عبارت، صفحه‌بندی، زمان‌بندی گزارش‌ها، RBAC کامل، پشتیبانی فرمت‌های بیشتر.

---

# منابع

- اسناد رسمی Go, Echo, GoAdmin, SQLite
- کد پروژه حاضر و کامنت‌های توضیحی


---

# پیوست A: بسط فصل اول (مبانی نظری و فناوری‌ها)

## A.1) زبان Go — تاریخچه، مدل هم‌زمانی و مدیریت حافظه

- تاریخچه مختصر: آغاز در 2009 توسط راب پایک، کن تامپسون و رابرت گریسمر در گوگل، الهام‌گرفته از C اما ساده‌تر و امن‌تر.
- مدل هم‌زمانی:
  - goroutine: نخ سبک‌وزن با زمان‌بند کاربر-فضا؛ راه‌اندازی ارزان و مناسب کارهای IO-bound.
  - channel: ارتباط امن بین goroutine ها؛ الگوی «حافظه را به اشتراک نگذارید، پیام را به اشتراک بگذارید».
  - select: چند-plexing روی چند کانال برای ساخت state machine های شفاف.
- مدیریت حافظه و GC:
  - garbage collector با تأخیر کم (low-latency)؛ مناسب سرویس‌های وب.
  - escape analysis و stack growth پویا؛ حافظه کمتر، عملکرد بهتر.
- مدیریت خطا (error handling):
  - برگشت مقادیر خطا کنار نتیجه؛ خوانایی بالا و کنترل دقیق مسیرهای خطا.
  - wrapping/annotating خطاها با fmt/errors برای بافت بیشتر.
- ابزارها (toolchain):
  - `go mod` برای مدیریت وابستگی‌ها؛ `go fmt` برای یکدست‌سازی؛ `go test` برای تست؛ `go vet` برای تحلیل ایستا.

## A.2) Echo — مسیرها، میان‌افزارها و بهترین‌عمل‌ها

- تعریف مسیرها با گروه‌بندی (Route Groups)، middleware سطح برنامه/گروه/مسیر.
- middleware های متداول: Recover, Logger, CORS, Gzip.
- نکات:
  - جداسازی هندلرها و DTO ها؛ اعتبارسنجی ورودی؛ تعیین دقیق `Content-Type`.
  - استفاده از context برای ارسال metadata و cancelation.

## A.3) GoAdmin — صفحه، جدول، فرم، منو و امنیت

- اجزای کلیدی:
  - Generator: تعریف صفحات برنامه از طریق تابع‌های سازنده (مثلاً `Tpage.Generators`).
  - Table/Form: افزودن ستون‌ها، فیلترها، گزینه‌های ویرایش درجا (inline edit) و SelectBox های پویا.
  - Menu: ایجاد آیتم‌های منو با آیکن‌ها و مسیرها (با `Tmenu`).
  - Theme: AdminLTE با اجزای UI آماده.
- امنیت:
  - کاربران/نقش‌ها/مجوزها با جداول `goadmin_*`؛ محدودسازی دسترسی‌ها از طریق slug و http path.
  - توصیه: فعال‌سازی احراز هویت، استفاده از HTTPS پشت reverse proxy.

## A.4) SQLite — تراکنش، هم‌زمانی، محدودیت‌ها

- مزایا: فایل‌محور، بدون سرویس جداگانه، مناسب نمونه‌سازی/محیط‌های کوچک.
- تراکنش‌ها: ACID با journal؛ برای نوشتن‌های هم‌زمان محدودیت دارد (Single-writer).
- محدودیت‌ها:
  - برای بار نوشتن بسیار بالا مناسب نیست؛ برای خواندن زیاد مناسب است.
  - اندازه فایل و قفل‌گذاری باید با care مدیریت شود.

## A.5) Front-end پایه — HTML/CSS/JS و نمودارها

- ساخت فرم‌های جستجو، جدول‌های نتایج، و داشبورد با ویجت‌های از پیش ساخته GoAdmin.
- ChartJS/ECharts برای نمایش روندها (تعداد خطاها در زمان، پراکندگی نوع رخدادها، ...).

---

# پیوست B: بسط فصل دوم (کد و معماری)

## B.1) راه‌اندازی برنامه — نمونه کدهای منتخب از main.go

```go
// تعریف پیکربندی GoAdmin و ثبت مسیرها
cfg := config.Config{ /* ... */ }
if err := eng.AddConfig(&cfg).
    AddGenerators(Tpage.Generators).
    AddDisplayFilterXssJsFilter().
    Use(e); err != nil {
    panic(err)
}

// مسیرهای کلیدی
eng.HTMLFile("GET", "/admin", "./html/helloDashboard.html", map[string]interface{}{})
eng.Data("GET", "/admin/healthCheck", healthCheck)
eng.Data("POST", "/search", DEL.SLHa.AdminHandler.SearchLogHandlerExternal)
```

توضیح:
- ثبت UI و API روی یک engine یکپارچه؛ استفاده از context داخلی GoAdmin برای سازگاری با تم و احراز هویت.

## B.2) Handler — تبدیل ورودی به DTO و خروجی JSON

```go
limit := 0
if strings.TrimSpace(c.FormValue("limit")) != "" {
    limit, _ = strconv.Atoi(c.FormValue("limit"))
}
req := params.SearchLogRequest{ /* year, month, day, ... */ }
if c.FormValue("countOnly") == "on" { req.CheckBox = true }

data, count, _ := Ha.adminSvc.GetFilesInFolder(req)
resp := params.SearchLogResponse{ Message: data, Count: fmt.Sprintf(" تعداد نتایج یافت شده:%d", count) }
c.Response.Header.Set("Content-Type", "application/json")
c.HTML(http.StatusOK, string(mustJSON(resp)))
```

نکات:
- برش‌دادن فضای خالی، تبدیل امن عدد، پشتیبانی از حالت شمارش-only.
- قراردادن `Content-Type` صحیح.

## B.3) Service — قرارداد و تزریق وابستگی

```go
type SearchLogs interface {
    GetFilesInFolder(params.SearchLogRequest) ([]string, int64, error)
}

type Service struct { SearchLogs SearchLogs }
```

نکته: جداسازی وابستگی امکان تست unit با mock را فراهم می‌کند.

## B.4) Search — مسیر فایل‌ها، پردازش موازی و فیلتر on-the-fly

```go
func (SL *SearchLogs) GetFileNames(filter params.SearchLogRequest) (int, error) {
    folder := base(SL.BasePath, filter) // logs/YYYY-MM-DD[/errors]
    filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
        // فیلتر پوشه errors بر اساس LogType و فیلتر نام فایل
        // افزودن به FilesPath و سپس ValidPath برای فایل‌های غیرخالی
        return nil
    })
    reverse(SL.FilesPath)
    return len(SL.FilesPath), nil
}
```

```go
func (SL *SearchLogs) ProcessFile(inProgress chan *backscanner.Scanner, out chan map[string]interface{}, wg *sync.WaitGroup) {
    wg1 := &sync.WaitGroup{}
    sem := make(chan struct{}, 10)
    for sc := range inProgress {
        sem <- struct{}{}
        wg1.Add(1)
        go func(scanner *backscanner.Scanner){
            defer func(){ wg1.Done(); <-sem }()
            SL.processFile(scanner, inProgress, out)
        }(sc)
    }
    wg1.Wait()
}
```

```go
func (SL *SearchLogs) CreateLine(line []byte, fileName string) map[string]interface{} {
    entry := map[string]interface{}{"fileName": fileName}
    if json.Valid(line) { /* parse JSON and time */ } else { /* استخراج JSON/زمان از متن */ }
    return entry
}
```

```go
func (SL *SearchLogs) ProcessLine2(in chan *backscanner.Scanner, out chan map[string]interface{}, isEnd *bool, req params.SearchLogRequest) ([]string, int64) {
    var keep []map[string]interface{}
    for {
        e := <-out
        if _, ok := e["eof"]; ok { /* مدیریت پایان همه فایل‌ها */ }
        if passFilter(req, JustformatEntry(e)) {
            keep = append(keep, e)
            if req.Limit != 0 && len(keep) == req.Limit { *isEnd = true; ClearInProgress(in); break }
        }
    }
    sort.Sort(ByTime(keep))
    return toJSONLines(keep), int64(len(keep))
}
```

## B.5) Backscanner — خواندن از انتها به ابتدا

```go
func (s *Scanner) LineBytes() (line []byte, pos int, err error) {
    if s.err != nil { return nil, 0, s.err }
    for {
        if i := bytes.LastIndexByte(s.buf, '\n'); i >= 0 {
            line, s.buf = dropCR(s.buf[i+1:]), s.buf[:i]
            return line, s.pos + i + 1, nil
        }
        s.readMore()
        if s.err != nil {
            if s.err == io.EOF && len(s.buf) > 0 { return dropCR(s.buf), 0, nil }
            return nil, 0, s.err
        }
    }
}
```

مزیت: خواندن مؤثر فایل‌های بزرگ بدون لود کامل در حافظه.

---

# پیوست C: بسط فصل سوم (کارایی، امنیت، تست)

## C.1) کارایی — پیچیدگی و تیونینگ

- پیچیدگی:
  - خواندن خطوط O(N) بر حسب تعداد خطوط بررسی‌شده.
  - مرتب‌سازی نهایی O(K log K) روی K خط منتخب.
- تیونینگ:
  - افزایش ظرفیت کانال‌ها برای throughput؛ تنظیم `DefultReadLine` برای تعادل latency/throughput.
  - محدودسازی goroutine ها با semaphore بر اساس منابع سیستم.
- سنجش:
  - زمان کل جستجو، تعداد فایل‌ها، اندازه فایل‌ها، نرخ سطر بر ثانیه.

## C.2) امنیت — ورودی، مجوز، استقرار

- ورودی:
  - sanitize و validate پارامترها؛ محدودسازی طول و کاراکترها.
  - جلوگیری از مسیرگردی (`..`/`/`).
- مجوز:
  - استفاده از نقش/مجوز GoAdmin برای محدودسازی مسیر `/search`.
- استقرار:
  - HTTPS پشت Nginx، تنظیم header های امنیتی، لاگ رویدادها.

## C.3) تست — سناریوهای پیشنهادی

- واحد:
  - `GetFileNames` با ترکیب‌های مختلف پارامترها.
  - `CreateLine` برای ورودی‌های JSON/متنی با/بدون زمان معتبر.
  - `FindWordInTexts2` برای ترکیبات include/exclude.
- ادغامی:
  - ساخت درخت پوشه لاگ با فایل‌های نمونه، فراخوانی `/search` و تأیید خروجی.
- بار/کارایی:
  - تولید فایل‌های بزرگ و سنجش latency و مصرف حافظه.

---

# پیوست D: راهنمای اجرا و استقرار پیشرفته

- پیش‌نیازها: Go 1.22، gcc (برای CGO)، دسترسی خواندن به مسیر لاگ.
- بیلد: `go build ./...` (خاموشی هشدار sqlite3: `CGO_CFLAGS="-Wno-return-local-addr"`).
- اجرا: `go run main.go -BasePath=""` و دسترسی به `http://localhost:1235/admin`.
- استقرار: باینری + systemd service + Nginx reverse proxy.
