import { notFound } from "next/navigation";
import React from "react";

// File: /src/app/dkmh/[dotId]/page.tsx

interface Subject {
  id: number;
  code: string;
  name: string;
  credits: number;
  instructor: string;
}

interface PageProps {
  params: {
    dotId: string;
  };
}

// fetch danh sách các môn đã đăng ký trong đợt
async function fetchRegisteredSubjects(dotId: string): Promise<Subject[]> {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_BASE_URL}/api/registrations?dotId=${dotId}`,
    { cache: "no-store" }
  );
  if (!res.ok) {
    throw new Error("Failed to fetch registered subjects");
  }
  return res.json();
}

export default async function Page({ params: { dotId } }: PageProps) {
  let subjects: Subject[];

  try {
    subjects = await fetchRegisteredSubjects(dotId);
  } catch (error) {
    // nếu không tìm thấy hoặc có lỗi
    return notFound();
  }

  return (
    <div style={{ padding: "1rem" }}>
      <h1>Danh sách môn đã đăng ký (Đợt {dotId})</h1>
      {subjects.length === 0 ? (
        <p>Chưa có môn nào được đăng ký.</p>
      ) : (
        <table border={1} cellPadding={8} cellSpacing={0}>
          <thead>
            <tr>
              <th>STT</th>
              <th>Mã học phần</th>
              <th>Tên học phần</th>
              <th>Tín chỉ</th>
              <th>Giảng viên</th>
            </tr>
          </thead>
          <tbody>
            {subjects.map((subj, idx) => (
              <tr key={subj.id}>
                <td>{idx + 1}</td>
                <td>{subj.code}</td>
                <td>{subj.name}</td>
                <td>{subj.credits}</td>
                <td>{subj.instructor}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
