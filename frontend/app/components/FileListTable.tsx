import { FileInfo } from "../models/file";

type FileListTableProps = {
  files: FileInfo[];
};

const FileListTable = ({ files }: FileListTableProps) => {
  return (
    <div className="overflow-x-auto">
      <table className="table">
        {/* head */}
        <thead>
          <tr>
            <th>Name</th>
            <th>Size</th>
            <th>Is Dir</th>
            <th>Last Modified</th>
          </tr>
        </thead>
        {/* body */}
        <tbody>
          {files.map((file, index) => (
            <tr key={index}>
              <td>{file.name}</td>
              <td>{file.size}</td>
              <td>{file.isDir ? "Directory" : "File"}</td>
              <td>{new Date(file.modTime).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default FileListTable;
