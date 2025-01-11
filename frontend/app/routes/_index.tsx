import { useLoaderData } from "@remix-run/react";
import FileListTable from "../components/FileListTable";
import { FileInfo } from "../models/file";

// TODO: Remove this hardcoded path
const testPath = "/mnt/verniy-share/Videos/";

export const loader = async () => {
  const params = new URLSearchParams({
    path: testPath,
  });

  const files = await fetch(
    `http://localhost:8080/api/file/list?${params.toString()}`,
    {
      method: "POST",
    }
  ).then((res) => res.json());

  const fileInfoArray: FileInfo[] = files.Metadata.map(
    (file: {
      name: string;
      size: number;
      is_dir: boolean;
      mod_time: string;
    }) => ({
      name: file.name,
      size: file.size,
      isDir: file.is_dir,
      modTime: file.mod_time,
    })
  );

  return { files: fileInfoArray };
};

export default function Index() {
  const { files } = useLoaderData<typeof loader>();

  return <div>{files && <FileListTable files={files} />}</div>;
}
