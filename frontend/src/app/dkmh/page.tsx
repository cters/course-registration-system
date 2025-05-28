type Student = {
  id: number;
  name: string;
  major: string;
  year: number;
};

const mockData: Student[] = [
  { id: 1, name: "Nguyen Van A", major: "Computer Science", year: 3 },
  { id: 2, name: "Tran Thi B", major: "Mathematics", year: 2 },
  { id: 3, name: "Le Van C", major: "Physics", year: 4 },
];

export default function DkmhPage() {
  return (
    <main className="p-8">
      <h1 className="text-2xl font-bold mb-4">Student List</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full border border-gray-200">
          <thead>
            <tr className="bg-gray-100">
              <th className="px-4 py-2 border">ID</th>
              <th className="px-4 py-2 border">Name</th>
              <th className="px-4 py-2 border">Major</th>
              <th className="px-4 py-2 border">Year</th>
            </tr>
          </thead>
          <tbody>
            {mockData.map((student) => (
              <tr key={student.id} className="hover:bg-gray-50">
                <td className="px-4 py-2 border">{student.id}</td>
                <td className="px-4 py-2 border">{student.name}</td>
                <td className="px-4 py-2 border">{student.major}</td>
                <td className="px-4 py-2 border">{student.year}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </main>
  );
}
