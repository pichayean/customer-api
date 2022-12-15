# GO vs CSharp
## web api performance tests.



# GO:
http_req_duration..............: avg=20.98ms min=9.95ms med=20.82ms max=63.34ms p(90)=26.6ms  p(95)=29.27ms
iterations............................: 3540   19.574038/s
http_reqs............................: 3540   19.574038/s
vus.....................................: 20     min=20      max=20
http_req_failed...................: 0.00%  ✓ 0         ✗ 3540
# CSharp:
http_req_duration..............: avg=192.78ms min=40.32ms med=77.19ms max=23.75s p(90)=122.76ms p(95)=143.97ms
iterations............................: 3306   16.271378/s
http_reqs..........................: 3306   16.271378/s
vus....................................: 20     min=20      max=20
http_req_failed..................: 0.36%  ✓ 12        ✗ 3294

# summary:
 vus => เท่ากัน ที่ 20 users
 http_req_duration => Go ทำได้ดีกว่า csharp percentile ที่ 95 Go ทำได้ 29.27ms | csharp ทำได้ 143.97ms
 iterations => Go ทำได้ดีกว่า Csharp Go ทำได้ 3540  คิดเป็น 19.574038/s | csharp ทำได้ 3306  คิดเป็น 16.271378/s
 http_req_failed => Go ทำได้ดีกว่า csharp Go failed 0.00%  [✓ 0 ,  ✗ 3540]  |  csharp [0.36%  ✓ 12, ✗ 3294]
Container Resource : CPU: 30, Ram: 318

- ✨Magic ✨
![alt text](https://github.com/pichayean/customer-api/blob/main/8350D280-3FBF-4EAE-AAF0-8D5AEB932FD0.jpeg?raw=true)
![alt text](https://github.com/pichayean/customer-api/blob/main/9F537FA9-E4F0-4D96-A81E-5FD6C28856ED.jpeg?raw=true)
