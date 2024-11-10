{ ID: 1, Task: "買い物に行く", Priority: 15, Status: false, Deadline: "2024-11-05" },
    { ID: 2, Task: "ジムで運動する", Priority: 3, Status: true, Deadline: "2024-11-10" },
    { ID: 3, Task: "本を読む", Priority: 0, Status: false, Deadline: "2024-11-05"},
    { ID: 4, Task: "部屋の掃除", Priority: 19, Status: false, Deadline: "2024-11-08" },
    { ID: 5, Task: "仕事のメールを返信", Priority: 1, Status: true, Deadline: "2024-11-05" },
    { ID: 6, Task: "友達とランチ", Priority: 0, Status: false, Deadline: "2024-11-15" },
    { ID: 7, Task: "映画を見る", Priority: 7, Status: false, Deadline: "2024-11-05" },
    { ID: 8, Task: "週末の予定を立てる", Priority: 12, Status: true, Deadline: "2024-11-12" },
    { ID: 9, Task: "歯医者の予約を取る", Priority: 8, Status: false, Deadline: "2024-11-05" },
    { ID: 10, Task: "レシピを考える", Priority: 20, Status: true, Deadline: "2024-11-20" },
    { ID: 11, Task: "日記を書く", Priority: 0, Status: false, Deadline: "2024-11-05" },
    { ID: 12, Task: "銀行に行く", Priority: 4, Status: false, Deadline: "2024-11-18" },
    { ID: 13, Task: "衣替えをする", Priority: 17, Status: true, Deadline: "2024-11-05" },
    { ID: 14, Task: "プレゼン資料を準備", Priority: 2, Status: false, Deadline: "2024-11-14" },
    { ID: 15, Task: "植物に水をやる", Priority: 0, Status: false, Deadline: "2024-11-05" },
    { ID: 16, Task: "ストレッチをする", Priority: 13, Status: true, Deadline: "2024-11-22" },
    { ID: 17, Task: "新しいレシピを試す", Priority: 5, Status: false, Deadline: "2024-11-05" },
    { ID: 18, Task: "オンラインコースを受講", Priority: 18, Status: false, Deadline: "2024-11-25" },
    { ID: 19, Task: "プロジェクトの進捗確認", Priority: 16, Status: true, Deadline: "2024-11-05" },
    { ID: 20, Task: "部屋の模様替えを考える", Priority: 6, Status: false, Deadline: "2024-11-30" },

    CREATE TABLE todoitems (
    id INT PRIMARY KEY AUTO_INCREMENT,
    task VARCHAR(255) NOT NULL,
    priority INT,
    status TINYINT(1),
    deadline VARCHAR(255)
);

INSERT INTO todoitems (id, task, priority, status, deadline) VALUES
(1, "買い物に行く", 15, 0, "2024-11-05"),
(2, "ジムで運動する", 3, 1, "2024-11-10"),
(3, "本を読む", 0, 0, "2024-11-05"),
(4, "部屋の掃除", 19, 0, "2024-11-08"),
(5, "仕事のメールを返信", 1, 1, "2024-11-05"),
(6, "友達とランチ", 0, 0, "2024-11-15"),
(7, "映画を見る", 7, 0, "2024-11-05"),
(8, "週末の予定を立てる", 12, 1, "2024-11-12"),
(9, "歯医者の予約を取る", 8, 0, "2024-11-05"),
(10, "レシピを考える", 20, 1, "2024-11-20"),
(11, "日記を書く", 0, 0, "2024-11-05"),
(12, "銀行に行く", 4, 0, "2024-11-18"),
(13, "衣替えをする", 17, 1, "2024-11-05"),
(14, "プレゼン資料を準備", 2, 0, "2024-11-14"),
(15, "植物に水をやる", 0, 0, "2024-11-05"),
(16, "ストレッチをする", 13, 1, "2024-11-22"),
(17, "新しいレシピを試す", 5, 0, "2024-11-05"),
(18, "オンラインコースを受講", 18, 0, "2024-11-25"),
(19, "プロジェクトの進捗確認", 16, 1, "2024-11-05"),
(20, "部屋の模様替えを考える", 6, 0, "2024-11-30");